#!/bin/bash
if [ $# -ne 1 ];
then
  echo "We need hostname"
  exit -1
fi
interval=150
hostname=$1 #104.154.100.74
port=":30513"
api_path=http://${hostname}${port}/api/v1/query


# docker stop Error
echo "stop docker $(date +%H):$(date +%M):$(date +%S)"
ssh -o StrictHostKeychecking=no -tt intern@${hostname} 'sudo systemctl stop docker' #stop
sleep $interval
ssh -o StrictHostKeychecking=no -tt intern@${hostname} 'sudo systemctl start docker' #start
echo "restart docker $(date +%H):$(date +%M):$(date +%S)"

# Check prometheus_health
prometheus_healthy=$(curl -i 'http://'${hostname}${port}'/-/healthy' 2>/dev/null | head -n 1 | cut -d$' ' -f2)
until [[ ${prometheus_healthy} = "200" ]]; do
  echo 'Prometheus is Not Healthy yet'
  prometheus_healthy=$(curl -i 'http://'${hostname}${port}'/-/healthy' 2>/dev/null | head -n 1 | cut -d$' ' -f2)
  sleep 5
done
# Recheck prometheus_ready for API
count_pod=$(curl -s ${api_path} --data-urlencode "query=count(kube_pod_info)" | jq -c '.data.result[].value[1]' | sed 's/"//g')
until [[ ${count_pod} != "" ]]; do
    echo 'Prometheus is Healthy But Not Ready'
    count_pod=$(curl -s ${api_path} --data-urlencode "query=count(kube_pod_info)" | jq -c '.data.result[].value[1]' | sed 's/"//g')
    sleep 5
done
echo "Now Prometheus is Healthy"
echo Count of Whole pods : $count_pod

# Array of Unhealthy pods (Error / Pending/ Crash ..)
query_json_of_condition_false_pods='(kube_pod_status_ready{condition="false"}==1)'
query_json_of_condition_false_pods_labeling_phase=''${query_json_of_condition_false_pods}' * on(pod) group_left(phase) label_replace( ( kube_pod_status_phase==1) , "phase" ,"$1", "condition", "(.+)" )'
query_json_of_succeded_pods='kube_pod_status_phase{phase="Succeeded"} ==1'

query_json_of_wrong_pods='sum by (pod, phase) ('${query_json_of_condition_false_pods_labeling_phase}')'\
'unless '\
'sum by(pod,phase) ( '${query_json_of_succeded_pods}' )'

wrong_pods=$(curl -s -g ${api_path} --data-urlencode "query=${query_json_of_wrong_pods}" | jq -c '.data.result[].metric.pod')
wrong_pod_array=( ${wrong_pods})
echo ${wrong_pods} | jq '.'
echo {wrong_pod_count : ${#wrong_pod_array[@]}}
init_count=${#wrong_pod_array[@]}


# Exception unhealthy pods
# for i in ${!wrong_pod_array[@]};do
#     if [[ ${wrong_pod_array[i]} =~ 'cap-elasticsearch-client' ]] || [[ ${wrong_pod_array[i]} =~ 'nvidia-device-plugin-daemonset' ]];then
#         unset wrong_pod_array[i]
#     fi
# done

# Check status of Unhealthy pods
index=0
until [ ${#wrong_pod_array[@]} -eq 0 ]; do
    pod_name=$(echo ${wrong_pod_array[${index}]}| sed 's/"//g')
    query_satus_of_wrong_pod='sum by (condition) (kube_pod_status_ready{pod="'${pod_name}'"}==1)'

    recheck_wrong_pod_status=$(curl -s -g ${api_path} --data-urlencode "query=${query_satus_of_wrong_pod}" | jq -c '.data.result[].metric.condition'| sed 's/"//g')
    echo [now-${#wrong_pod_array[@]}][${index}/${init_count}]: ${pod_name} : ${recheck_wrong_pod_status}

    if [[ ${recheck_wrong_pod_status} = "true" ]]; then
        unset wrong_pod_array[${index}]
        echo "          Pod_Recovery :  ${pod_name} -> now pod_count : ${#wrong_pod_array[@]}}"
    fi
    if [ ${index} -eq $((${init_count} - 1)) ]; then
        index=0
        echo Loop
        sleep 10
    else
        index=$((${index}+1))

    fi
done
echo "All Pod status: Running (Recovery Succeeded) "
# wait for loki recovery
sleep 60