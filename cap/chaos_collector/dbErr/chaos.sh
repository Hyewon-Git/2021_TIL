#!/bin/bash
if [ $# -ne 1 ];
then
  echo "We need kubeconfig directory"
  exit -1
fi
kubeconfig=$1

interval=3m

katib_db_err="katib-mysql"
pipeline_db_err='mysql'
metadata_db_err='metadata-db'
db_err_array=(${katib_db_err} ${pipeline_db_err} ${metadata_db_err})
db_cnt=${#db_err_array[@]}

random_db_err=${db_err_array[$(($RANDOM% $db_cnt))]}
#run_err
echo do error $random_db_err
kubectl --kubeconfig $kubeconfig scale deployment -n kubeflow $random_db_err --replicas 0
sleep $interval
#recover
echo do recover $random_db_err
kubectl --kubeconfig $kubeconfig scale deployment -n kubeflow $random_db_err --replicas 1
