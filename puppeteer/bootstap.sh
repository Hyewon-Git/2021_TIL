#!/bin/bash
# CAP harbor의 aiops 프로젝트에 먼저 로그인 되어 있어야 함.
if [ $# -ne 1 ];
then
  echo "Please enter the image version"
  exit -1
fi
version=$1

sudo docker build -t cap-activity-automation:$version .
sudo docker tag cap-activity-automation:$version cap.dudaji.com:31480/aiops/cap-activity-automation:$version
sudo docker push cap.dudaji.com:31480/aiops/cap-activity-automation:$version
