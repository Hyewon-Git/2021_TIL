# Helm chart for daiops
## Synopsis
AIOps의 여러 repository의 yaml파일들을 `daiops` helm chart 로 생성하여 복잡한 어플리케이션의 배포를 관리한다. 

values.yaml파일에 지정된 템플릿화 되어있는 chart의 변수를 통해, 사용자가 새로운 설정값을 재설정하여 배포할 수 있다.

<br>

## QuickStart

**1. daiops namespace 생성**

`kubectl create namesapce [namespace 명]`

`kubectl create namespace daiops`

<br/>

**2-1. helm chart 기본값으로 설치**

`helm install [chart 명] [chart 폴더 경로] [옵션]`

`helm install daiops . --namespace daiops`
		
namespace 추가하여 지정하지 않으면 --> default namespace에 설치

<br/>

**2-2. helm chart 커스터마이징하여 설치**

--values (또는 -f): 오버라이드(override)할 YAML 파일을 지정한다. 여러 번 지정할 수 있지만 가장 오른쪽에 있는 파일이 우선시된다.

--set: 명령줄 상에서 오버라이드(override)를 지정한다.
설치시 다음과 같이 옵션을 추가해주면 된다.
    
`--set key=value`

<br/>

**daiops chart 구조**

: 4가지 repository( `alt`, `backend`, `modelServer`, `web` )의 yaml파일을 helm chart로 만들었다.
```
- charts
    : 해당 디렉토리에는 종속성을 가지고 있는 다른 helm chart를 저장한다.
    
        - influxdb/
        - minio/
        - postgresql/
        - rabiitmq/

- template
    : 생성하고자 하는 kubernetes resource형태로 폴더가 있고, 해당 폴더 내부에 AIOps repository명칭을 딴 yaml파일이 존재한다.
        (web repository는 frontend로 명칭 사용)

        - deployments/
        - services/
        - helpers.tpl : template manifest파일들에서 공유하는 항목 정의 (차트 전반에 걸쳐 재사용하는 항목)

- values.yaml
    : 커스터마이징 가능 변수들의 집합
      해당 변수를 오버라이딩, 삭제 가능하다.
```

<br/><br/>

## Reference
**Helm**

: kubernetes에서 애플리케이션을 배포, 관리하면 복잡성 up & 배포시간 up

→ Helm을 이용해 애플리케이션을 보다 빠르게 배포 관리 가능

https://helm.sh/ko/docs/helm/helm/

	* Helm 특징 

    1. 복잡한 어플리케이션 배포 관리
    2. Hooks
        Kubernetes 환경에서 helm 차트로 설치, 업그레이드,삭제 그리고 롤백과 같은 애플리케이션 생명주기의 개입할 수 있는 기능을 Hook을 통하여 제공

    3. 릴리즈 관리
        Helm으로 배포된 애플리케이션은 하나의 릴리즈로 불립니다. 해당 릴리즈는 배포된 애플리케이션의 버전 관리를 가능하도록 합니다.

    
