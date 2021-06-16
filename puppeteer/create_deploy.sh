#!/bin/bash
if [ $# -lt 1 ];
then
  echo "We need the target server ip. > create_deploy.sh 0.0.0.0 "
  exit -1
fi

server_ip=$1
server_str=$(echo ${server_ip} | sed -e "s/\./\-/g" )

echo server_ip ${servier_ip}
echo server_str ${server_str}

cat <<EOF > deployment-${server_str}.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: aiops-pupeteer-${server_str}
  namespace: aiops 
  labels:
    app: aiops-pupeteer-${server_str}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: aiops-pupeteer-${server_str}
  template:
    metadata:
      labels:
        app: aiops-pupeteer-${server_str}
    spec:
      containers:
        - name: puppeteer
          image: cap.dudaji.com:31480/aiops/cap-activity-automation:v0.0.1
          imagePullPolicy: Always
          ports:
            - containerPort: 80
          env:
            - name: TARGET_SERVER
              value: "http://${servier_ip}:31380"
EOF

kubectl apply -f deployment-${server_str}.yaml
