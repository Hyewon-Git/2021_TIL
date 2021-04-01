# Dudaji ICT Internship

git 참조하기 :[https://ychae-leah.tistory.com/55](https://ychae-leah.tistory.com/55)

(2021/3/2~)

> 21.03.02(화) ~03(수)

- zsh

    :: oh-my-zsh 설치

    ::플러그인 설치

    ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled.png)

    [https://medium.com/harrythegreat/zsh와-함께-사용할-플러그인-추천-6가지-8f9b8b7f3c24](https://medium.com/harrythegreat/zsh%EC%99%80-%ED%95%A8%EA%BB%98-%EC%82%AC%EC%9A%A9%ED%95%A0-%ED%94%8C%EB%9F%AC%EA%B7%B8%EC%9D%B8-%EC%B6%94%EC%B2%9C-6%EA%B0%80%EC%A7%80-8f9b8b7f3c24)

    -Auto Suggestions : 터미널 histroy 기반 단어 추천

    전체 제안 Command를 적용하려면 → or End

    단어 단위로 제안 Command를 적용하려면 ⌥ + →

    -Syntax Highlighting

    zsh에서 해당 Command가 실행될 수 있는지 또는 에러인지 확인할 수 있습니다.

    [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting) 플러그인을 설치해보겠습니다.

    -OSX

    맥 유저에게 정말 유용한 기능을 제공합니다. 간단히 명령어만 소개해드리면

    - tab — 새로운 탭 열기
    - ofd — 현재 디렉터리를 파인더에서 열기(open . 와 유사).
    - pfd — 가장 최근의 파인더의 PATH
    - pfs = 현재 파인더 디렉터리 주소값(여러개라면 복수개 표시)
    - cdf = 현재 파인더로 디렉터리 이동하기
    - pushdf = 파인더에 파일 넣기

    -bat

- docker for mac
    1. docker-desktop 설치
    2. docker app실행 후 kubernetes 설정
    3. [https://brunch.co.kr/@sokoban/81](https://brunch.co.kr/@sokoban/81)

        kubernetes  설치 후 

    [https://www.44bits.io/ko/post/news--release-docker-desktop-with-kubernetes-to-stable-channel](https://www.44bits.io/ko/post/news--release-docker-desktop-with-kubernetes-to-stable-channel)

    ?? docker kubernetes 개념 좀더 자세히 알아보기!!!

    [https://brenden.tistory.com/109](https://brenden.tistory.com/109)

    ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%201.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%201.png)

- mac 사용법

    : app 다운 및 작동법 익히기!!!

    - zsh : bash 를 대신하는 shell환경 ( 확장기능 플러그인기능 많)
    - iTerms2

        : 맥의 기본 Terminal을 대체할 수 있는 터미널 에뮬레이터입니다. 화면분할 등 여러가지 기능, 테마 변경 등 다양한 기능

         + color 설정

         +  ZSH_THEME = agnoster로 : git의 상태를 알려줌!  폰트 do2...

        [https://rap0d.github.io/tip/2020/06/15/mac_oh_my_zsh/](https://rap0d.github.io/tip/2020/06/15/mac_oh_my_zsh/)

    - Spectacle : 화면 분할
    - VScode : MS에서 개발한 에디터

[리눅스 명령어](https://www.notion.so/277fb2309da14de7ad08317ccfe6f71f)

- 코드스타일

    : 코딩을 편리하게 공유하기위해 아래같은걸 씀!  ex) pep 8 같은거(style guide)

    -formatter : 띄어씌기 마침표같은거

    -linter : 내용적 (이게좀더 큼! 교집합인 부분이 많음)

    - dudaji에서는

        python - formatter: black 커스터마이즈가 잘안됨!

        vscode 터미널에서 : pip install balck( 설치부터)

        black [main.py](http://main.py) —> 이러면 formatter가 적용됨

        >> 매번 할수없으니깐 

        폴더에 setting.json 파일만들어서  git코드적으면 —→ save 누르면 자동으로 포메팅되게 (.vscode  폴더안에 파일 만들어줄수있음 / 개인설정에서하면 전체적 적용에 할수있음! settings에

         pip install  pylint

> 03.04(목)

- [x]  11:30 까지 쿠버네티스 정리 및 Nginx 이미지 띄우기

[https://www.oss.kr/info_techtip/show/7524a1f3-00f7-4182-b8da-e1ba67899a60](https://www.oss.kr/info_techtip/show/7524a1f3-00f7-4182-b8da-e1ba67899a60)

[https://kubernetes.io/docs/concepts/services-networking/service/](https://kubernetes.io/docs/concepts/services-networking/service/)

[https://kubernetes.io/ko/docs/tutorials/services/source-ip/](https://kubernetes.io/ko/docs/tutorials/services/source-ip/)

[https://kubernetes.io/docs/concepts/workloads/controllers/deployment/](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

#.쿠버네티스 자동완성 플러그인 : [https://kubernetes.io/ko/docs/tasks/tools/install-kubectl/](https://kubernetes.io/ko/docs/tasks/tools/install-kubectl/)

[쿠버네티스 이해하기](https://www.notion.so/47d4c6935b7a4310834a7e83a9215887)

   > go로 짠 백엔드 - 쿠버네티스(여러서버 aws, k8s,dudaji등) 연동

gpu brokerage 플랫폼의 backend 개발을 맡기로 했고 go와 kubernetes에 익숙해지기 위해 2주정도 사전준비 후 개발 진행할 예정입니다. 사전준비 기간에는 golang 학습과 k8s sdk를 활용해 local k8s를 제어하는 간단한 api 서버를 만들어보려고 합니다.

[ go언어 학습 & 프로그래머스] + [ go 로 api 서버 ] > [ 쿠버네티스   Task 수행] 

교재 : [http://pyrasis.com/go.html](http://pyrasis.com/go.html)
         [http://golang.site/go](http://golang.site/go/b)
학습공간 : [https://play.golang.org/](https://play.golang.org/)

> 03.05 (금)~08(월)09(화)10(수)11(목)12(금)

- [ ]  vscode 사용법 익히기

[go  언어학습 [참조: 가장 빨리만나는  go 언어 ] ](https://www.notion.so/go-go-3b5ee646fe984feb9ccb6f36e5f599f1)

[vi 편집기 명령어 익히기](https://www.notion.so/vi-f72f26994ae6469c939f9988e4d9f59d)

++ clean코드읽어보기

반복문속에서 조금이라도 덜 반복되고 return으로빼주는게좋은가

반복문끝나고 밖에서 깔끔하게 가독성좋게 적는게 좋은가

[Golang-programmers](https://www.notion.so/Golang-programmers-e8d5b6f8e24848aeb1f8c7058e42ea2e)

> 3.14 (일) 자료구조

[자료구조](https://www.notion.so/82ed4a0bff834b0cbea1e7fe48b69934)

> 3.15(월) go 간단한 웹서버 만들기

아주 간단 웹서버기능 - net/http 패키지 // 복잡,다양한 기능을 요하는 웹서버인 경우 web framework사용

- Go Web Framework TYPE

    [Untitled](https://www.notion.so/21fa425e074d4139869af87ac0a6fb4e)

    [https://sarc.io/index.php/miscellaneous/2033-golang-rest-api](https://sarc.io/index.php/miscellaneous/2033-golang-rest-api)

    [https://streamls.tistory.com/77](https://streamls.tistory.com/77) >> gin framework

- api 서버란

    [https://velog.io/@za1013/API-서버란-무엇일까](https://velog.io/@za1013/API-%EC%84%9C%EB%B2%84%EB%9E%80-%EB%AC%B4%EC%97%87%EC%9D%BC%EA%B9%8C)

- go VScode  제대로 설치

    $ GOPATH :  기본 defaul폴더에 < pkg다운받으면 여기안으로 들어가짐!

    vscode 에서 내가 사용하고싶은 폴더에서

    .vscode 폴더에 >> settings.json

    ```json
    {
      "go.lintOnSave": true,             // 파일 저장시 linter 실행 
      "go.lintTool": "golint",           // liter tool로 golint 사용 
      "go.formatTool": "goreturns",      // formatter로 goreturns 사용 
      "editor.formatOnSave": true,       // 파일 저장시 formatter 실행 
    }
    ```

    go.mod << 터미널서 `go mod init 현재디렉토리(workspace폴더명)`  명령어로 다음 파일 생성해줌 ($ GOPATH에 있는 것들 모듈로 연동)

    main.go 작성 >> 터미널서 `go run main.go` 를해주면 터미널에서 실행됨!

    [https://blog.golang.org/using-go-module](https://blog.golang.org/using-go-module) (원본)

    [https://johngrib.github.io/wiki/golang-mod/](https://johngrib.github.io/wiki/golang-mod/)   (번역본)

    [https://docs.github.com/en/github/working-with-github-pages](https://docs.github.com/en/github/working-with-github-pages)

- 여러 예제
    - 카카오챗봇 만들기 [https://choiseokwon.tistory.com/292](https://choiseokwon.tistory.com/292)
    - 

[Gin Web](https://www.notion.so/Gin-Web-617c2d46c4a141b79998811c190d0141)

 > 참고 튜토리얼 

[https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)

[https://streamls.tistory.com/244?category=830512](https://streamls.tistory.com/244?category=830512)

[https://gin-gonic.com/ko-kr/docs/examples/bind-single-binary-with-template/](https://gin-gonic.com/ko-kr/docs/examples/bind-single-binary-with-template/)

[https://github.com/gin-gonic](https://github.com/gin-gonic)

- Beego Web

    [https://beego.me/docs/quickstart/controller.md](https://beego.me/docs/quickstart/controller.md)

    [https://sourcegraph.com/github.com/kubernetes/client-go@abf396d787a7442759dd7d9315005bedc639e134/-/blob/examples/dynamic-create-update-delete-deployment/main.go#L167](https://sourcegraph.com/github.com/kubernetes/client-go@abf396d787a7442759dd7d9315005bedc639e134/-/blob/examples/dynamic-create-update-delete-deployment/main.go#L167)v

> 3.19(금) 쿠버네티스 pod-env / Go API 서버

- Pod yaml파일에서  env 설정
    - 1. string으로

        ```yaml
        apiVersion: v1
        kind: Pod
        metadata:
          name: envar-demo
          labels:
            purpose: demonstrate-envars
        spec:
          containers:
          - name: envar-demo-container
            image: gcr.io/google-samples/node-hello:1.0
        */  command: [ "sh", "-c"]
            args:
            - while true; do
                  echo -en '\n';
                  printenv MY_NODE_NAME MY_POD_NAME MY_POD_NAMESPACE;
                  printenv MY_POD_IP MY_POD_SERVICE_ACCOUNT;
                  sleep 10;
                done;
        */ 이렇게 넣어줄수도 있다 ! 실행하면 바로뜨게! 
            env:
            - name: DEMO_GREETING
              value: "Hello from the environment"
            - name: DEMO_FAREWELL
              value: "Such a sweet sorrow"
        ```

        >> 실행방법

        1. `kubectl exec [ pod이름] -- printenv`  다음 파드를 실행시킨후 — *명령어를 실행해라
        2. `kubectl exec -it [pod이름] /bin/bash`  >> `echo $[환경변수이름]`
        3.  */ ~ /* command를 넣어서   `kubectl logs [pod이름]` 을 해주면 바로확인가능
    - 2. secret으로
        - secret  생성방법
            1. yaml파일로 생성

                ```yaml
                apiVersion: v1
                kind: Secret
                metadata:
                  name: test-secret
                data:
                  username: bXktYXBw
                  password: Mzk1MjgkdmRnN0pi
                ```

                `kubectl apply -f [yaml파일명]`

            2. kubectl 명령어로 직접 시크릿 생성

                `kubectl create secret generic [secret명] --from-literal='[key명]=[value값 ]' --from-literal='[또다른 key명]=[value값]'`

        1. volume을 통해  secret 데이터에 접근가능한 pod생성
        2. secret데이터를 이용하여 컨테이너 환경변수 env정의
            1.  secret 생성 

                `kubectl create secret generic backend-user --from-literal=backend-username='backend-admin'`

            2. yaml파일을 통해 pod 생성
                1. pod 의 yaml파일에 미리 정의된 secret명을 pod의 사양인 환경변수에 할당

                    ```yaml
                    apiVersion: v1
                    kind: Pod
                    metadata:
                      name: env-single-secret
                    spec:
                      containers:
                      - name: envars-test-container
                        image: nginx
                        env:
                        - name: SECRET_USERNAME
                          valueFrom:
                            secretKeyRef:
                              name: backend-user   ///secret명!!!!!
                              key: backend-username
                    ```

                2. secret의 모든 key-value쌍을 컨테이너 환경변수로

                    :envForm을 사용하여 모든 secret의 데이터를 컨테이너 환경 변수로 정의

                    ```yaml
                    apiVersion: v1
                    kind: Pod
                    metadata:
                    	name: envfrom-secret
                    spec:
                    	containers:
                    	  - name: envars-test-container
                    			image: nginx
                    			envFrom:
                    			    -secretRef:
                    							name: test-secret
                    ```

        >> 실행방법

         kubectl apply -f [yaml파일명 : ~~~.yml]

        kubectl get po [pod명] -o wide : 파드의  IP주소와 파드가 배포된  노드 표시

        `kubectl exec -i -t [pod명] -- /bin/sh -c 'echo $[secret명]'`

        [https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/](https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/)

    - 3. ConfigMap으로
        - Configmap생성하기
            1. create 명령어로    `kubectl create configmap <map-name> <data-source>`
                - diretory 자체를

                    `kubectl create configmap [configmap명] --from-file=[configmap파일이 있는 폴더 경로로]`

                - file 로

                    `kubectl create configmap [configmap명] --from-file=[configmap파일]  --from-file=[configmap파일] 등등등`  여러configmap파일을  하나의 configamp으로 생성할때 다음과같이 나열하면됨!

                - file에서 configmap을 만들때 사용할 키 정의

                    `kubectl create configmap [configmap명] --from-file=[kye 이름]=[configmap파일 경로/파일명]`

                - 리터럴 값에서 configmap 생성

                    `kubectl create configmap [configmap명] --from-literal=[key명]=[value] --from-literal=[또다른key명]=[value]`

            2. 생성기에서 Configmap만들기 >> yaml파일로
                - file로 : yaml파일에다가 configmap파일 경로/파일명 적어주기!!

                    ```yaml
                    cat <<EOF >./kustomization.yaml
                    configMapGenerator:
                    - name: game-config-4
                      files:
                      - configure-pod-container/configmap/game.properties
                    EOF
                    ```

                    `kubectl apply -k . [경로]/[configmap명] created`

                - 리터럴로

                    ```yaml
                    # Create a kustomization.yaml file with ConfigMapGenerator
                    cat <<EOF >./kustomization.yaml
                    configMapGenerator:
                    - name: special-config-2
                      literals:
                      - special.how=very
                      - special.type=charm
                    EO
                    ```

        컨피그맵을 사용하여 파드 내부에 컨테이너를 구성할 수 있는 네 가지 방법이 있다.

        1. 컨테이너 커맨드와 인수 내에서
        2. 컨테이너에 대한 환경 변수
            1. configmap 미리 생성
            2. configmap에 미리 정의된 key값을 pod의 환경변수에 할당>> pod생성을 위한 yaml파일
                1. configMapKeyRef:

                    ```yaml
                    apiVersion: v1
                    kind: Pod
                    metadata:
                      name: dapi-test-pod
                    spec:
                      containers:
                        - name: test-container
                          image: k8s.gcr.io/busybox
                          command: [ "/bin/sh", "-c", "env" ]
                          env:
                            # Define the environment variable
                            - name: SPECIAL_LEVEL_KEY #환경변수이름
                              valueFrom:
                                configMapKeyRef:
                                  # The ConfigMap containing the value you want to assign to SPECIAL_LEVEL_KEY
                                  name: special-config #사용하는 configmap 이름
                                  # Specify the key associated with the value
                                  key: special.how #사용할 configmap내에서 key값 
                      restartPolicy: Never
                    ```

                2. envForm을 이용해 configmap 내의 모든 데이터를 컨테이너 환경수로 정의

                    ```yaml
                    apiVersion: v1
                    kind: Pod
                    metadata:
                      name: dapi-test-pod
                    spec:
                      containers:
                        - name: test-container
                          image: k8s.gcr.io/busybox
                          command: [ "/bin/sh", "-c", "env" ]
                          envFrom:
                          - configMapRef:
                              name: special-config
                      restartPolicy: Never
                    ```

            3. `kubectl create -f [yaml파일명]`  : yaml파일대로 pod생성

        3. 애플리케이션이 읽을 수 있도록 읽기 전용 볼륨에 파일 추가
        4. 쿠버네티스 API를 사용하여 컨피그맵을 읽는 파드 내에서 실행할 코드 작성

            == 컨피그맵과 데이터를 읽기 위해 코드를 작성해야 한다는 것을 의미

        [https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#create-a-configmap-using-kubectl-create-configmap](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#create-a-configmap-using-kubectl-create-configmap)

 yaml파일 형식으로 보고싶을때 : `k get [종류 pod,node,secrets,configmaps 등등] [이름] -o yaml`

- Volume
    - 볼**륨**

        도커는 다소 느슨하고, 덜 관리되지만 [볼륨](https://docs.docker.com/storage/)이라는 개념을 가지고 있다. 도커 볼륨은 디스크에 있는 디렉터리이거나 다른 컨테이너에 있다. 도커는 볼륨 드라이버를 제공하지만, 기능이 다소 제한된다.

        쿠버네티스는 다양한 유형의 볼륨을 지원한다. [파드](https://kubernetes.io/ko/docs/concepts/workloads/pods/)는 여러 볼륨 유형을 동시에 사용할 수 있다. 임시 볼륨 유형은 파드의 수명을 갖지만, 퍼시스턴트 볼륨은 파드의 수명을 넘어 존재한다. 결과적으로, 볼륨은 파드 내에서 실행되는 모든 컨테이너보다 오래 지속되며, 컨테이너를 다시 시작해도 데이터가 보존된다. 파드가 더 이상 존재하지 않으면, 볼륨은 삭제된다.

        기본적으로 볼륨은 디렉터리일 뿐이며, 일부 데이터가 있을 수 있으며, 파드 내 컨테이너에서 접근할 수 있다. 디렉터리의 생성 방식, 이를 지원하는 매체와 내용은 사용된 특정 볼륨의 유형에 따라 결정된다.

        볼륨을 사용하려면, `.spec.volumes` 에서 파드에 제공할 볼륨을 지정하고 `.spec.containers[*].volumeMounts` 의 컨테이너에 해당 볼륨을 마운트할 위치를 선언한다.

        컨테이너의 프로세스는 도커 이미지와 볼륨으로 구성된 파일시스템 뷰를 본다. [도커 이미지](https://docs.docker.com/userguide/dockerimages/)는 파일시스템 계층의 루트에 있다. 볼륨은 이미지 내에 지정된 경로에 마운트된다. 볼륨은 다른 볼륨에 마운트할 수 없거나 다른 볼륨에 대한 하드 링크를 가질 수 없다. 파드 구성의 각 컨테이너는 각 볼륨을 마운트할 위치를 독립적으로 지정해야 한다.

    - **hostPath**

        `hostPath` 볼륨은 호스트 노드의 파일시스템에 있는 파일이나 디렉터리를 파드에 마운트 한다. 이것은 대부분의 파드들이 필요한 것은 아니지만, 일부 애플리케이션에 강력한 탈출구를 제공한다.

        예를 들어, `hostPath` 의 일부 용도는 다음과 같다.

        - 도커 내부에 접근할 필요가 있는 실행중인 컨테이너. `/var/lib/docker` 를 `hostPath` 로 이용함
        - 컨테이너에서 cAdvisor의 실행. `/sys` 를 `hostPath` 로 이용함
        - 파드는 주어진 `hostPath` 를 파드가 실행되기 이전에 있어야 하거나, 생성해야 하는지 그리고 존재해야 하는 대상을 지정할 수 있도록 허용함

        필요한 `path` 속성 외에도, `hostPath` 볼륨에 대한 `type` 을 마음대로 지정할 수 있다.

        필드가 `type` 에 지원되는 값은 다음과 같다.

        ```yaml
        apiVersion: v1
        kind: Pod
        metadata:
          name: test-pd
        spec:
          containers:
          - image: k8s.gcr.io/test-webserver
            name: test-container
            volumeMounts:
            - mountPath: /test-pd
              name: test-volume
          volumes:
          - name: test-volume
            hostPath:
              # 호스트의 디렉터리 위치
              path: /data
              # 이 필드는 선택 사항이다
              type: Directory
        ```

    - **컨피그맵(configMap)**

        컨피그맵은 구성 데이터를 파드에 주입하는 방법을 제공한다. 컨피그맵에 저장된 데이터는 `configMap` 유형의 볼륨에서 참조되고 그런 다음에 파드에서 실행되는 컨테이너화된 애플리케이션이 소비한다.

        컨피그맵을 참조할 때, 볼륨에 컨피그맵의 이름을 제공한다. 컨피그맵의 특정 항목에 사용할 경로를 사용자 정의할 수 있다. 다음 구성은 `log-config` 컨피그맵을 `configmap-pod` 라 부르는 파드에 마운트하는 방법을 보여준다.

        ```yaml
        apiVersion: v1
        kind: Pod
        metadata:
          name: configmap-pod
        spec:
          containers:
            - name: test
              image: busybox
              volumeMounts:
                - name: config-vol
                  mountPath: /etc/config
          volumes:
            - name: config-vol
              configMap:
                name: log-config
                items:
                  - key: log_level
                    path: log_level
        ```

        log-config 컨피그맵은 볼륨으로 마운트되며, log_level 항목에 저장된 모든 컨텐츠는 파드의 /etc/config/log_level 경로에 마운트된다. 이 경로는 볼륨의 mountPath 와 log_level 로 키가 지정된 path 에서 파생된다.

        >> completed — image에서 command로 명령어로 다실행되고 끝난상태의미 ! 

        실행 명령어는 log에서 확인할수있다!!!

[Go API 서버 docker에 올리기](https://www.notion.so/Go-API-docker-4b90f0516b374b9dba76e1f3c1ba5942)

```
kubectl config view# 병합된 kubeconfig 설정을 표시한다.# 동시에 여러 kubeconfig 파일을 사용하고 병합된 구성을 확인한다KUBECONFIG=~/.kube/config:~/.kube/kubconfig2

kubectl config view

# e2e 사용자의 암호를 확인한다
kubectl config view -o jsonpath='{.users[?(@.name == "e2e")].user.password}'

kubectl config view -o jsonpath='{.users[].name}'# 첫 번째 사용자 출력
kubectl config view -o jsonpath='{.users[*].name}'# 사용자 리스트 조회
kubectl config get-contexts# 컨텍스트 리스트 출력
kubectl config current-context# 현재 컨텍스트 출력
kubectl config use-context my-cluster-name# my-cluster-name를 기본 컨텍스트로 설정# 기본 인증을 지원하는 새로운 사용자를 kubeconf에 추가한다
kubectl config set-credentials kubeuser/foo.kubernetes.com --username=kubeuser --password=kubepassword

# 해당 컨텍스트에서 모든 후속 kubectl 커맨드에 대한 네임스페이스를 영구적으로 저장한다
kubectl config set-context --current --namespace=ggckad-s2

# 특정 사용자와 네임스페이스를 사용하는 컨텍스트 설정
kubectl config set-context gce --user=cluster-admin --namespace=foo\
  && kubectl config use-context gce

kubectl config unset users.foo# foo 사용자 삭제
```

> 3.29 (월) ~ 4.1(목) https /  Istio

- HTTPS 로 API Server 생성
    - https  개념
        - **HTTP** : Hyper Text Transfer **Protocol**

            Hyper Text를 전송하기 위해 만든 프로토콜- Hyper Text (ex: HTML)

            Client (웹 브라우저) ←→Server사이의 요청/응답 프로토콜입니다.

            웹상에서 통신을 할 때 사용

            [http://naver.com](http://naver.com/) : "naver.com"와 통신 시 HTTP 프로토콜을 활용

            - HTTP의 약점
                1. **암호화하지 않은 통신이기 때문에 도청할 수 있다.**

                    통신 경로 상에 있는 네트워크 기기나 케이블이나 컴퓨터 등을 전부 자기 자신이 소유할 수는 없습니다. 악의를 가진 사람이 기기를 통해 통신 내용을 도청할 수 있습니다.

                2. **통신 상대를 확인하지 않기 때문에 위장할 수 있다.**
                    - (Request의 출처) Client를 확신할 수 없다.
                        - HTTP는 누가 Request를 보내도 Response하는 구조입니다. 신원이 보장된 특정 Client와만 통신할 수 없습니다.
                        - 또한 대량의 Request를 통한 Dos 공격의 위험이 있습니다.
                    - (Response의 출처) Server를 확신할 수 없다.
                        - Response를 보낸 Server가 내가 의도한 Server인지 확신할 수 있습니다. 위장 Server라는 위험성이 있습니다.
                3. **완전성을 증명할 수 없기 때문에 변조 할 수 있다.**

                    완전성: 정보의 정확성

                    Client와 Server가 보낸 정보를 중간에 누군가 바꿀 위험성이 있습니다.

        ![https://user-images.githubusercontent.com/37133536/88912829-da733000-d29a-11ea-8b1d-c14643eca410.png](https://user-images.githubusercontent.com/37133536/88912829-da733000-d29a-11ea-8b1d-c14643eca410.png)

        - HTTPS ( **HTTP** over **S**ecure Socket Layer )

            :증명서를 통해 서버 또는 클라이언트의 신원을 확인하고, 데이터를 암호화, 인증, 안정성 보호를 할 수 있는 프로토콜

            SSL,TLS 프로토콜 위에서 실행되는 HTTP를 의미

            **HTTPS는 SSL(Secure Socket Layer)을 이용한 HTTP 통신 방식**입니다.

            HTTP에 SSL의 껍질을 씌운 것

            HTTPS는 HTTP와 별개인 새로운 프로토콜이 아닙니다.

            HTTP는 보안에 취약하기 때문에 HTTP에 다른 보안 프로토콜을 조합한 것

            HTTP 통신을 하는 소켓 부분을 SSL이나 TLS 프로토콜로 대체합니다

            - HTTP의 보안이 강화된 버전이다.
            - SSL 프로토콜을 통해 세션 데이터를 암호화 한다.
            - 사용자 컴퓨터와 방문한 사이트 간에 전송되는 사용자 데이터의 무결성과 기밀성을 유지할 수 있게 해주는 인터넷 통신 프로토콜이다.
            - 데이터암호화 (도청 / 추적 / 도용에 대한 보호)무결성 (변조 / 손상 방지)인증 (요청에 대한 신뢰 보장)
                - ssl인증서 - 목적
                    - 클라이언트가 접속한 서버가 신뢰 할 수 있는 서버임을 보장한다.
                    - SSL 통신에 사용할 공개키를 클라이언트에게 제공한다.
                    - 즉, CA로부터 신뢰할 수 있는 사이트라는 것을 인증받으면 공개키를 얻을 수 있고, 공개키를 통해 데이터를 암호화, 복호화 한다

            - Https를 쓰는 이유

                네트워크를 통해 패킷을 주고 받음  << 여러 공격자로 부터 위협

                 : 특정 서버에 요청을 주고 받을 때 수 많은 라우터와 스위치를 거침

                이때 누군가 우리의 패킷을 훔칠수도 (Sniffing) 있음

                ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%202.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%202.png)

                그래서 네트워크 데이터를 암호화시킴! ← 데이터 유출을 막고자

                오늘날 가장 널리 쓰이는 암호화 방식 = SSL/TLS1 - 인증서 서명 방식

                                                : 1. 인증서(신뢰할수있는사람인지)-검증작업 —> 2.  데이터 암호화

                >> 인증서를 사용하지 않은 웹서버 : 데이터가 암화화가 X → 데이터 유출 가능

                1. 패킷 스니핑(Server - Client 간 통신을 중간에서 훔쳐보는것) 방지
                2. 파밍 (가짜 홈페이지를 활용해 사용자 금융정보 빼돌리는 기법) 방지

                [https://webactually.com/2018/11/16/http에서-https로-전환하기-위한-완벽-가이드/](https://webactually.com/2018/11/16/http%EC%97%90%EC%84%9C-https%EB%A1%9C-%EC%A0%84%ED%99%98%ED%95%98%EA%B8%B0-%EC%9C%84%ED%95%9C-%EC%99%84%EB%B2%BD-%EA%B0%80%EC%9D%B4%EB%93%9C/)

                - **기밀성** HTTPS는 인터넷과 같은 공공 매체에서 두 참여자 간의 통신을 보호한다. 예를 들어, HTTPS가 없다면 와이파이 액세스 포인트를 운영하는 사람은 액세스 포인트를 사용하는 사람이 온라인에서 무언가를 구입할 때 신용카드와 같은 개인정보를 볼 수도 있다.
                - **무결성** HTTPS는 변조되지 않은 정보로 목적지에 도달하게 한다. 예를 들어, 와이파이가 웹사이트에 광고를 추가하거나, 대역폭을 절약하고자 이미지 품질을 저하시키거나, 읽는 기사의 내용을 변조할 수 있지만 HTTPS는 웹사이트를 변조할 수 없도록 한다.
                - **인증** HTTPS를 통해 웹사이트의 진위 여부를 확인할 수 있다. 예를 들어, 와이파이 액세스 포인트을 운영하는 사람이 가짜 웹사이트를 브라우저에 보낼 수도 있다. HTTPS는 `example.com`이라는 웹사이트가 실제로 `example.com`인지 확인한다. 일부 인증서는 `yourbank.com`이 YourBank.Inc라는 걸 알리기 위해 해당 웹사이트의 법적 신원을 검사하기도 한다.

                    >>> 가짜 사이트로 되는 걸 방지 >> EX) 위장 롯데리아 사이트

        - SSL (Secure Scoket Layer)

            HTTP와 독립된 프로토콜 (다방면에 사용)

            SSL == TLS 같은 말

            대칭키를 이용 암호화 통신 —> 암호화 통신을 위해서 **SSL인증서**필요!!

            - SSL 인증서

                : 누군가 이 사이트가 신뢰있는 사이트라고 인증해주기윟나 인증서

                ( 인증서소유자email,이름, 인증서용도, 유효기간, 발행기관,pubic key 등이 포함)

            - SSL을 활용한 통신 방법

                Client Server 통신에 앞서 Server는 **CA에서 인증서를 받는습니다.**

                - CA (Certificate authority)

                    : **공인된 기관**에서 Server가 믿을 수 있는 서버인지 보증하는 SSL 보증서를 발급합니다.

                     — CA기관마다 보안 강도가 다르기도 합니다.

                     — 뒤에 실습에서 살펴볼 인증서는 보안 강도가 약한 수준이지만, 보안 수준이 높은 인증서는 사업자등록증까지 요구한다고 합니다.

                    자체 CA (사설 CA)로도 SSL 인증서를 발급할 수 있습니다.

                    - 사설 CA도 HTTPS 통신입니다. 하지만 브라우저 입장에서는 안전하지 않다고 판단합니다.
            - 통신 방법(구체적)
                1. Client가 Server에 최초 접속하면서 2가지 정보를 보냄

                    random dataClient가 생성한 random data입니다.이건 나중에 사용됩니다.암호화 기법 목록을 보냅니다.SSL에서 사용되는 암호화 기법은 여러가지입니다.

                    ![https://user-images.githubusercontent.com/37133536/88912853-e232d480-d29a-11ea-99e9-298d2d594ba0.png](https://user-images.githubusercontent.com/37133536/88912853-e232d480-d29a-11ea-99e9-298d2d594ba0.png)

                2. Server는 3가지 정보를 보냅니다.

                    ![https://user-images.githubusercontent.com/37133536/88912864-e65ef200-d29a-11ea-8c6a-ff19e3f20ead.png](https://user-images.githubusercontent.com/37133536/88912864-e65ef200-d29a-11ea-8c6a-ff19e3f20ead.png)

                    - random data : Server가 생성된 random data
                    - Client가 보낸 암호화 기법 중 자신도 사용할 수 있고, 가장 안정된 암호화 기법
                        - 사용할 암호화기법
                    - 인증서
                        - 인증서에는 서비스 정보와 public key를 보냅니다.
                        - 서비스 정보: 인증서를 발급한 CA, 서비스의 도메인 등등
                            - 서비스 정보는 private key로 암호화된 상태입니다.

                3.  Server가 신뢰할 수 있는 서버인지 '인증서'를 통해 검증

                - ca 인증서 구조

                    [https://m.blog.naver.com/alice_k106/221468341565](https://m.blog.naver.com/alice_k106/221468341565)

                    ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%203.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%203.png)

                    공인 CA기관의 공개키는 컴퓨터에 이미 설치되어있음 (보통)

                    — Mac OS 경우 : 키체인에 

                    - Chain of Trust :

                        ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%204.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%204.png)

                        (1) Chain of Trust의 원리에 의해 하위 인증서가 신뢰할 수 있는지를 알 수 있으며, (2) 하위 인증서의 내용물이 변조되었는지를 알 수 있게 된다.

                        ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%205.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%205.png)

                (1) Client는 믿을 수 있는 CA에서 발급한 인증서인지 확인합니다.

                ![https://user-images.githubusercontent.com/37133536/88912882-eb23a600-d29a-11ea-8ab2-c83d8d5094ce.png](https://user-images.githubusercontent.com/37133536/88912882-eb23a600-d29a-11ea-8ab2-c83d8d5094ce.png)

                - Browser는 믿을 수 있다고 판단한 CA 기관 목록을 가지고 있습니다.
                - 공인 CA (믿을 수 있는 CA) / 사설 CA에서 발급한 인증서는 각각 다른 형태로 표시됩니다.

                (2).  Client는 실제 CA 기관에서 발급한 인증서인지 확인합니다.

                ![https://user-images.githubusercontent.com/37133536/88912903-f4ad0e00-d29a-11ea-820f-587cbe7f6137.png](https://user-images.githubusercontent.com/37133536/88912903-f4ad0e00-d29a-11ea-820f-587cbe7f6137.png)

                - Server가 보낸 인증서에는 서비스 정보와 public key가 있습니다.

                    private key와 public key는 하나의 쌍을 이룹니다.

                - private key로 암호화된 서비스를 public key로 복호화할 수 있다면,  ==⇒ "전자서명"
                    - 두 key는 pair입니다.
                    - CA 기관에서 발급한 인증서라고 할 수 있습니다.
                    - **즉, 내가 기대한 서버인지 확인하게 됩니다.**

                ⇒ **비대칭키 방식** (public, private key를 사용하는 암호화 방식)

                - 대칭키 비대칭키
                    - 대칭키 : 암/복호화 키가 같음

                        - 비공개키 사용

                    - 비대칭키 = 공개키 암호 : 암/복호화 키가 다름

                        - 공개키/비공개키 둘 다 사용

                1. 통신에 사용할 key를 Client와 Server는 공유합니다. ⇒ **공통키 방식** (하나의 key를 공유하는 암호화 방식)

                    ![https://user-images.githubusercontent.com/37133536/88912934-0393c080-d29b-11ea-82ed-d3e7b197398c.png](https://user-images.githubusercontent.com/37133536/88912934-0393c080-d29b-11ea-82ed-d3e7b197398c.png)

                - premaster secret은 앞서 언급한 random data를 합쳐 생성합니다.
                - public key로 암호화하면 private key로만 복호화가능 —> premaster secret은 안전
                - premaster secret → master secret → session key

                     ⇒ 일련의 과정을 거쳐 client와 server는 공통키를 가지게 됩니다.

                ![https://user-images.githubusercontent.com/37133536/88912954-0d1d2880-d29b-11ea-9c06-6dc7c0da445c.png](https://user-images.githubusercontent.com/37133536/88912954-0d1d2880-d29b-11ea-9c06-6dc7c0da445c.png)

                - 공통의 session key로 데이터를 암호화/복호화할 수 있습니다.

                ⇒ **대칭키** ( session key로 암호화/복호화 )

            ![https://media.vlpt.us/images/moonyoung/post/5102b861-7201-49e6-a97c-e519288c9e16/image.png](https://media.vlpt.us/images/moonyoung/post/5102b861-7201-49e6-a97c-e519288c9e16/image.png)

             >> 대칭키를 공유할 때 사용하는 암호화 기법 = 공개키방식

             >> 실 데이터를 공유할 때 사용하는 암호화 기법 = 대칭키 방식

            - 왜 이런 방식을 취할까? = **하이브리드 암호 시스템**

                주목해야 할 점은 대칭키 (1개의 키) 방식을 취하되, 그 키를 공유할 때 공개키 (public key, private key) 방식을 취한다는 점입니다.

                - 대칭키 : 기밀성 유지한 통신 가능 But 키 배송 문제
                    - 한 개의 key로 암호화 복호화
                    - 속도가 빠르다.
                    - 탈취의 위험이 있다.
                    - 암호를 주고 받는 사람들 사이에 대칭키를 전달하는 것이 어렵다는 점이다. 대칭키가 유출되면 키를 획득한 공격자는 암호의 내용을 복호화 할 수 있기 때문에 암호가 무용지물이 되기 때문이다.
                - 비대칭키 = 공개키 (공개키, 비밀키) : 키배송문제가 없음 But 대칭키보다 느림, 중간자 공격에 약함
                    - 공개키로 암호화, 비밀키로 복호화
                    - **속도가 느리다.**
                    - 인증의 기능까지 제공한다.
                    - 

                    공개키 방식은 두개의 키를 갖게 되는데 A키로 암호화를 하면 B키로 복호화 할 수 있고, B키로 암호화하면 A키로 복호화 할 수 있는 방식이다. 이 방식에 착안해서 두개의 키 중 하나를 비공개키(private key, 개인키, 비밀키라고도 부른다)로하고, 나머지를 공개키(public key)로 지정한다. 비공개키는 자신만이 가지고 있고, 공개키를 타인에게 제공한다. 공개키를 제공 받은 타인은 공개키를 이용해서 정보를 암호화한다. 암호화한 정보를 비공개키를 가지고 있는 사람에게 전송한다. 비공개키의 소유자는 이 키를 이용해서 암호화된 정보를 복호화 한다. 이 과정에서 공개키가 유출된다고해도 비공개키를 모르면 정보를 복호화 할 수 없기 때문에 안전하다. 공개키로는 암호화는 할 수 있지만 복호화는 할 수 없기 때문이다.

                    이 방식은 이렇게 응용할 수도 있다. 비공개키의 소유자는 비공개키를 이용해서 정보를 암호화 한 후에 공개키와 함께 암호화된 정보를 전송한다. 정보와 공개키를 획득한 사람은 공개키를 이용해서 암호화된 정보를 복호화 한다. 이 과정에서 공개키가 유출된다면 의도하지 않은 공격자에 의해서 데이터가 복호화 될 위험이 있다. 이런 위험에도 불구하고 비공개키를 이용해서 암호화를 하는 이유는 무엇일까? 그것은 이것이 데이터를 보호하는 것이 목적이 아니기 때문이다. 암호화된 데이터를 공개키를 가지고 복호화 할 수 있다는 것은 그 데이터가 공개키와 쌍을 이루는 비공개키에 의해서 암호화 되었다는 것을 의미한다. 즉 공개키가 데이터를 제공한 사람의 신원을 보장해주게 되는 것이다. 이러한 것을 전자 서명이라고 부른다.

                **대칭키의 단점 (1개의 키를 공유하면서 해킹당할 수 있는 위험), 공개키의 단점 (속도가 느리고, 많은 컴퓨팅 파워가 필요) 모두를 보완하는 방법입니다.**

    - golang API server -https로 연결

        [http://webs.co.kr/?mid=https&page=1&listStyle=list&document_srl=3321270](http://webs.co.kr/?mid=https&page=1&listStyle=list&document_srl=3321270)

        - key 생성

            [https://blusky10.tistory.com/352](https://blusky10.tistory.com/352)

            1. private 키 생성

            `openssl genrsa -out [private.key] 2048`  : 비번없이

            `openssl genrsa -des3 -out 키이름 2048` : 비번 있이

            2. public 키 생성  = private키랑 쌍이 되는 공개키

            `openssl rsa -in [private.key] -pubout -out [public.key]`

            3. CSR (certificate signing request -인증서 서명 요청) 생성

            : ssl인증의 정보를 암호화하여 인증기관에 보내어 인증서를 발급받게하는 신청서

            ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%206.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%206.png)

            `openssl req -new -key [private.key] -out [private.csr]`

            4. CRT (Certificate) 인증서 

            : CSR을 생성했다면 CRT를 그냥 만들수있지만, 

            [https://namjackson.tistory.com/24](https://namjackson.tistory.com/24)

            - CA인증 원리

                [https://m.blog.naver.com/alice_k106/221468341565](https://m.blog.naver.com/alice_k106/221468341565)

            - self Signed로 만든 ca인증서로 인증까지 받은  server인증서 만들기

                ##.  사설 ca인증을 받아야 local에서 server가능

                1. rootCA.key

                    `openssl genrsa <암호화 알고리즘> -out 키이름 2048`

                    -aes256 

                2. rootCA 사설 CSR생성

                    `openssl req -x509 -new -nodes -key rootCA.key -days 3650 -out rootCA.pem`

                    ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%207.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%207.png)

                3. CRT생성

                    : 2에서 만들었던 csr을 나만의 커스텀 ca인 rootCA의 인증을 받아 private.crt로 생성

                    `openssl x509 -req -in [private.csr] -CA [rootCA.pem] -CAkey [rootCA.key] -CAcreateserial -out [private.crt] -days 3650`

            - self Signed 인증서 만드는 법
                1. CSR을 명시적으로 넣어서 인증서 만드는 방법

                `openssl x509 -req -days <유효날수> -in <인증사인요청파일.csr> -signkey <개인키.key> -out <인증서 파일명.crt>`

                2. CSR을 넣지 않고 암묵적으로 인증서 만드는 방법

                 `openssl req -new -x509 -days <유효날수> -key <개인키> -out <인증서파일명>`

            - 기타 openssl 설정사항( 인증서 확인/ 형식전환)
                - `openssl x509 -text -noout -in [인증서파일]` : 인증서내용확인가능
                - `openssl x509 -in [crt파일] -out [crt를pem파일명] -outform PEM`  : CRT파일을 PEM파일로 변환
                - 

                ```
                # key 변경
                openssl rsa -in server.key -text > private.pem
                # crt 변경
                openssl x509 -inform PEM -in server.crt > public.pem
                ```

                [https://www.letmecompile.com/certificate-file-format-extensions-comparison/](https://www.letmecompile.com/certificate-file-format-extensions-comparison/)

                [https://onestepcloser.tistory.com/142](https://onestepcloser.tistory.com/142)

            5. localhost로 접근시 chrome설정사항

            [https://velog.io/@jereint20/bypass-sslerrorpage](https://velog.io/@jereint20/bypass-sslerrorpage)

- Istio
    - 마이크로 서비스 구조 : 너무 잘게 나뉨 >>  circuit breaker패턴 적용

        마이크로 서비스 아키텍처에서 각 마이크로 서비스는 간단한 작업을 소유하며, REST API 요청과 같은 경량 통신 방식을 사용하여 클라이언트 또는 기타 마이크로 서비스와 통신

        ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%208.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%208.png)

        ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%209.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%209.png)

    - 서비스 매쉬

        이런 마이크로 서비스 문제를 풀기 위해 

        - 소프트웨어 계층이 아니라 인프라 측면에서 풀기 위해서 "**proxy**"이용

            ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%2010.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%2010.png)

             ++ proxy가 circuit breaker같은 역할

            proxy : "대리" 의미 = 중계 (보안상/ 시간절약(캐시)/병목현상방지)

            [https://brownbears.tistory.com/191](https://brownbears.tistory.com/191)

        각 프록시에 대한 설정 정보를 중앙 집중화된 컨트롤러가 통제하는 구조로

        ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%2011.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%2011.png)

        - **Data Plane** : 각 프록시들로 이루어져서 트래픽을 설정값에 따라 트래픽을 컨트롤 하는 부분
        - **Control Plane**: 데이타 플레인의 프록시 설정값들을 저장하고, 프록시들에 설정값을 전달하는 컨트롤러 역할을 하는 부분
    - Istio : 서비스 매쉬의 구조를 구현한 예

        : 조직에서 분산형 마이크로서비스 기반 앱을 어디서나 실행할 수 있도록 지원하는 오픈소스 서비스 메시

         = Envoy proxy를 Data Plane으로 사용하고 이를 컨트롤 해주는 오픈 소스 솔루션

        ![Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%2012.png](Dudaji%20ICT%20Internship%2036b1ead1d1854b47bd5d6f87a46ed93f/Untitled%2012.png)

        - Data plane : envoy를 서비스 옆에 붙여서 사이드카 형식으로 배포하여 서비스로 출입하는 트래픽을 envoy를 통해 통제

            (envoy : service discovery =  서비스 호출시 상대 서비스IP알아야함  —> control plane의 pilot에서 end-point참고)

        - Control Plane : data plane에 배포된 envoy를 컨트롤하는 부분
            - Pilot : envoy에 대한 설정 관리
            - Mixer :  access control, 정책 통제, 모니커링 지표 수집
            - Citadel : 보안 관련- 서비스 사용을 위한 사용자 인증
        - 기능
            - 트래픽 통제
                - 트래픽 분할 : 서로 다른 버전의 서비스 배포해놓고, 버전별 트래픽양 조절 기능 (카날리 테스트)
                - 컨텐츠 기반의 트래픽 분할 : 네트워크 패킷 내용 기반으로 라우팅 가능
            - 서비스간 안정성 제공 (Resilience)
                - 헬스체크 및 서비스 디스커버리

                    : pilot은 대상 서비스가 여러개의 인스턴스로 구성되있으며 이를 로드 밸런싱, 주기적, 핼스체크, 장애 서비스 자동 제거

            - 보안
                - 통신보안
                - 서비스 인증과 인가
                    - 서비스 간 인증
                    - 서비스 - 사용자간 인증
                    - 인가를 통한 권한 통제 (Authorization)
            - 모니터링

                네트워크 트래픽을 모니터링 > 서비스간 호출관계/ 서비스 응답시간/ 처리량등의 다양한 지표 수집하여 모니터링가능 > **Mixer**이용

            ⇒ Istio를 통해 조직은 마이크로서비스를 보호, 연결, 모니터링할 수 있으므로 엔터프라이즈 앱을 더욱 빠르고 안전하게 현대화할 수 있습니다.

            Istio는 기존의 분산형 애플리케이션에 투명하게 레이어링하여 배포 복잡성을 완화합니다.

    - Istio  설치 및 BookInfo예제

> 4.1(목) ~ OIDC

- OIDC란

    :
