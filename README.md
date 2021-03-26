# 2021_TIL
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
