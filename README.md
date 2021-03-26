# 2021_TIL
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
