# Puppeteer for cap-activity-automation
## Synopsis
AIOps의 log 수집 중 정상상태의 log에 대하여 마치 사용자의 활발한 CAP 활동 상황을 인위적으로 부여하기 위해 `Puppeteer`를 사용하여 CAP 활동을 만들어주었다. 
환경변수를 통해 활동을 부여하고자하는 서버를 지정해주면 된다.

## QuickStart
1. cap-activity-automation 빌드 및 cap harbor에 이미지 푸시
    ```(shell)
    ./bootstrap.sh <image version>
    ex) ./bootstrap.sh v0.0.1
    ```  

2. deployment 생성 및 배포 : create_deploy.sh을 실행하면 매개변수로 전달되는 TARGET_IP의 값을 기준으로 deployment 가 생성이 되며, aiops ns에 배포된다.
   ```(shell)
   > create_deploy.sh TARGET_IP

   ```
## Reference
* Puppeteer

    Puppeteer는 DevTools 프로토콜을 통해 Chrome 또는 Chromium을 제어하는 고급 API를 제공하는 노드 라이브러리

    브라우저에서 수동으로 수행 할 수있는 대부분의 작업은 Puppeteer를 사용하여 수행 가능

    https://github.com/puppeteer/puppeteer/tree/v9.0.0
