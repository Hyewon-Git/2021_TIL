# 🔐 SSO 🔐

**📌   목표 - Gitlab을 Keycloak SAML방식을 이용한 SSO 환경 구축**

  - 구축 환경 

    docker-desktop (mac) & kubernetes

    **local**에서 작동
    
    [1. Keycloak pod 띄우기](#keycloak-설치)

    [2. gitlab pod 띄우기](#gitlab-설치)

    [3. Keycloak -gitlab SSO 환경구축](#keycloak---gitlab-연동)
     (SAML 방식이용)
     
    [4. Keycloak Login, Register Customizing](#keycloak-기타-설정)
    
    [5. Keycloak API 이용](#keycloak-기타-설정)
---

## Keycloak 설치

local에서 작동하므로 ingress 없이 

`kubectl create -f https://raw.githubusercontent.com/keycloak/keycloak-quickstarts/latest/kubernetes-examples/keycloak.yaml`

두가지 방법 중 선택

1. `kubectl edit svc keycloak` 해서 loadbalancer 를 NodePort로 바꿔주던가

    Keycloak 페이지 주소 : `http://localhost:<nodeport>`
    
2. kecloack **설치 이전**에 (keycloak을 먼저설치하면 pending상태로 지속됨)

    nginx 컨테이너 nodeport로 expose
    
    --> 저절로  loadbalancer 의 External_IP가  localhost로 설정됨
    
    ~~어떻게 이렇게 작동된지는 모르겠음!! 원래 local에서는 loadbalancer실행 X~~
    
    Keycloak 페이지 주소 : http://localhost:8080

---
## Gitlab 설치

두가지 방법 중 선택 (여기서는 2번방법 이용)

1.  helm 사용

    설치 :[https://docs.gitlab.com/charts/installation/](https://docs.gitlab.com/charts/installation/)

2. image 이용해서 deployment > service : NodePort 로 설정
    - command 이용하는 방법

        `kubectl create deployment gitlab --image=gitlab/gitlab-ce:latest --port=80`

        `kubectl expose deployment gitlab --type=NodePort --port=80 --target-port=80 --name=gitlab-service`

    - yaml파일을 이용하는 방법 (command 방법은 nodeport를 지정해줄수 없다!)
        - gitlab.yaml

            ```yaml
            # gitlab.yaml
            apiVersion: apps/v1
            kind: Deployment
            metadata:
              name: gitlab
              labels:
                app: gitlab
            spec:
              replicas: 1
              selector:
                matchLabels:
                  app: gitlab
              template:
                metadata:
                  labels:
                    app: gitlab
                spec:
                  containers:
                  - name: gitlab-container
                    image: gitlab/gitlab-ce:latest
            ---
            apiVersion: v1
            kind: Service
            metadata:
                name: gitlab-service
            spec:
                ports:
                - nodePort: 30010 # Random assignment from 30000 to 32767, if omitted
                    port: 80 # Allocate the same value as targetPort if omitted
                    targetPort: 80
              selector:
                app: gitlab
            ```

        gitlab은 실행된 이후 5분정도 이후에 접속가능

         바로 직후엔  ~~curl: (52) Empty reply from server이거나  502 whoops something went wrong on our end 에러뜸~~

> gitlab설치 후 잘 안돌아가면 —> kubernetes  문제임 
> .  >> docker-desktop의 setting — resource > memory를 올려주어라!

---
## Keycloak - Gitlab 연동

Keycloak은 2가지 SAML, OpenId-connect 방식이 있다.

Gitlab은 여러  종류 OmniAuth provider와 통합한다.  keycloak은 SAML, openid-connect 방식 등의 OmniAuth provider(표준 identity provider protocol)을 제공한다.

여기에서는 SAML방식을 이용해서 SSO를 구현해보겠다.

### Keycloak setting

[http://localhost:8080/auth](http://localhost:8080/auth) 로
처음 접근시: [master realm = admin]계정으로 접근 (`아이디 admin/ 비번 admin` : keycloak설치시 yaml파일에서 초기 아이디,비번 지정해줌)

1. realm 생성

    : 초기에는 Master Realm만 존재

     but, 우리는 Master가 아닌 MSA에서 사용할 Realm이 필요하기때문에 생성

    ![img/Untitled.png](img/Untitled.png)

    "Add realm"로 생성

    Realm Settings 에서 
    { 일반설정, 로그인 화면 설정, 키 암호화 설정, 이메일 서비스 설정, 로그인/회원가입 등의 테마 설정, 캐시 설정, 토큰 설정} 등 —— > 필요시 공식문서 참고할것!
    ~~회원가입 및 로그인 , user field  추가는 뒤에서~~

2. user 생성

    realms 안에서 관리자를 통해 유저를 생성할 수 있다.

    <사용하고자하는 realm>환경임을 확인 -> Manage -> Users 클릭 -> Add User 클릭한다

    ![img/Untitled%201.png](img/Untitled%201.png)

    이후 생성되면 "credentials"탭에서 유저에게 임시비밀번호 발급해준다 (Set Password)해줌

3. client 생성
    gitlab 과 SAML 로 SSO 를 할꺼니깐 설정해줘야함
    
    - client 생성 시 "SAML"  Protocol로 생성 >> 이후 아래와 같이 설정

      ![img/Untitled%202.png](img/Untitled%202.png)

      ![img/Untitled%203.png](img/Untitled%203.png)

      ![img/Untitled%204.png](img/Untitled%204.png)

      IDP initiated SSO URL Name = client명과 같게 해주기

      해당 옵션 별 의미 : [https://www.keycloak.org/docs/latest/server_admin/#saml-clients](https://www.keycloak.org/docs/latest/server_admin/#saml-clients)


    - 해당 client의 Mapper 설정해주기

        MapperType > name,email,first_name,last_name만 기본 User property

        Property > User생성시 각 속성명을 매칭

        SAML Attribute Name > gitlab으로 전송되는 data명칭

        ![img/Untitled%205.png](img/Untitled%205.png)
        
         —> 오타 위에 saml attribute name : last_name으로 할것!

        keycloak의 data를 gitlab(SP: Service Provider)이 필요로 하는 값으로 Mapping하여 정확히 전송할 수있다. (metadata, role 등)

- **(번외) keycloak account페이지에서 로그인 및 계정 확인방법**

    `http://localhost:8080/auth/realms/<사용하고자하는 realm>/account` 

    (관리자페이지)생성한 realm에서 생성한 user정보로 로그인해서 계정 확인가능

      
### Gitlab setting

Pod로 접속하여 KeyCloak의 설정값을 바꿔줌

- 공식페이지 [https://docs.gitlab.com/ee/integration/omniauth.html](https://docs.gitlab.com/ee/integration/omniauth.html)
                  [https://docs.gitlab.com/ee/integration/saml.html](https://docs.gitlab.com/ee/integration/saml.html)
- gitlab의 설정파일을 열어서 편집

    `sudo vim /etc/gitlab/gitlab.rb`

    OmniAuth 공급자 설정

    공식페이지 [https://docs.gitlab.com/ee/administration/auth/oidc.html](https://docs.gitlab.com/ee/administration/auth/oidc.html)

    - **인증서 정보**

        Realm Settings > Keys> Public Keys의  Certificate정보 (클릭하여 값 복사하기)

        ![img/Untitled%206.png](img/Untitled%206.png)

    -  **OmniAuth 공급자 중 "SAML" 사용*

         [https://www.debyum.com/configure-gitlab-with-keycloak/](https://www.debyum.com/configure-gitlab-with-keycloak/)

        [https://edenmal.moe/post/2018/GitLab-Keycloak-SAML-2-0-OmniAuth-Provider](https://edenmal.moe/post/2018/GitLab-Keycloak-SAML-2-0-OmniAuth-Provider/?__cf_chl_captcha_tk__=abd99e96a0b76ba86f74686fd3ca96783d87b8d8-1618449110-0-Ac7yEOdq9cMmxXtZ05mffkm-WEt8xo_NfputD_f27Rqzh4jehkrLXH_jkbLwf7bbJrjdOEY8aJpa2IGwBfeWITecHQiJBG2N1taIvwrSsk9uvPv4EZP76YucqgQyQq-6mXR5x8hp2UG-ODDV-lT26fwpCBFqk3robXjakxTOShSjjSJOLXczyuc5KjOIJfQrCtrc4tsYquXSoUts953hqNyBPC1qObV5emUjzKYG8KULaYfpkYdji8vtX1cZNCPDpd7eKKjJRey6V4Ce8icUqaDf3bVIdMGmbHMBGTcuwRBFJZW5VI-feIe7m7UvbFmQZ0uV2yVqw9CDLOJAtlxwgVS8ZDht1hHM2Znt7dP8sDCOJ7BxkDYVccsjTK5bsg61w5qFZ8eBr3Xrfpw1KpfMtMXeAgp_VUSKeORiC4d2wwb9d0MXsXIqt1-cYb3xUUgocThJkLnPp3FaHktPV_Z-OMUye9UQmWI3_h-fAHzIbKR5lOCg0W6bHYhxfKv0dfCkwc0Lq46jSTAYbB2yxW5mXufairq7ABdKz29Cg1hH_giXL5n6dUGrBgTUgD7l6NqLFkWc9MBjOCoTfmZwt1VgExMDZEiYPsbObcCXgCJpMi-Q2A20D7byhu47SHNg5rbwlg8aIq5KD7nY_MID1qc2aCT7xKa4hON1vgf_86uG-nCW)

        [https://gist.github.com/int128/ab5839d5f59829840a0204f0c8cd8a8b](https://gist.github.com/int128/ab5839d5f59829840a0204f0c8cd8a8b)

        ```ruby
        gitlab_rails['omniauth_enabled'] = true

        #사용자가 먼저 계정을 수동으로 만들지 않고도 SAML을 사용하여 가입
        gitlab_rails['omniauth_allow_single_sign_on'] = ['saml']
        # gitlab_rails['omniauth_auto_sign_in_with_provider'] = 'saml'
        gitlab_rails['omniauth_block_auto_created_users'] = false

        # gitlab_rails['omniauth_auto_link_ldap_user'] = false

        #이메일 주소가 일치하는 경우 SAML 사용자를 기존 GitLab 사용자와 자동으로 연결
        gitlab_rails['omniauth_auto_link_saml_user'] = true 

        gitlab_rails['omniauth_providers'] = [
            {
              name: 'saml',
              args: {
                       assertion_consumer_service_url: 'http://localhost:30010/users/auth/saml/callback',
                       idp_cert:  "-----BEGIN CERTIFICATE-----\n<...인증서정보 복사해서 붙여넣기....>\n-----END CERTIFICATE-----\n",
                       idp_sso_target_url: 'http://localhost:8080/auth/realms/dudaji-200/protocol/saml/clients/gitlab',
                       issuer: 'gitlab',
                       name_identifier_format: 'urn:oasis:names:tc:SAML:2.0:nameid-format:persistent'
                     },
              label: 'Keycloak Login'
            }
        ]
        ```

    설정파일을 변경 후 무조건 설정변경 명령어를 해줘야함 `gitlab-ctl reconfigure`

    - log 확인 
    
        `cat /var/log/gitlab/gitlab-rails/production.log`

***

## Keycloak Login, Register Customizing

### **Keycloak 회원가입 페이지 설정**

회원가입 창을 띄우고 싶은 realm 선택  Realm Settings >  Login

 ![img/Untitled%207.png](img/Untitled%207.png)
 
**User registration - ON** : 회원가입 페이지 존재하게

(위에서의 기본설정에서는 off 되어있음)


### **Keycloak 로그인 페이지 꾸미기**

[https://www.keycloak.org/docs/latest/server_development/#scripts](https://www.keycloak.org/docs/latest/server_development/#scripts)

[https://www.baeldung.com/spring-keycloak-custom-themes](https://www.baeldung.com/spring-keycloak-custom-themes)

1. kubernetes 위에 있는 Pod ( Keycloak container)로 진입

     ****‼️ keycloak container 내부에서는 vim이나 nano가 없다**

          **→  volume 마운트를 이용해 파일을 생성하거나 수정해라!**

     ❗ 기존 테마인 base, keycloak, keycloak2.0 폴더 안에서 수정하는 것 보다 

    해당 폴더를 복사하여 새로운 폴더(새로운 테마)를 생성하여 수정 혹은 생성하는것이 좋다!

2. `/opt/jboss/keycloak/standalone/configuration/standalone.xml` 해당 파일의 값을

    ```xml
    <theme>
    	<staticMaxAge>-1</staticMaxAge>
    	<cacheThemes>false</cacheThemes>
    	<cacheTemplates>false</cacheTemplates>
    	...
    </theme>
    ```

    다음과 같이 바꿔라 → 개발단계에서 서버를 다시 시작하지 않고도 변경사항의 효과를 즉시 확인

3. keycloak의 theme 폴더로 진입

    #. Keycloak의 기본페이지 `$JBOSS_HOME` ⇒ /opt/jboss/keycloak

    `cd /opt/jboss/keycloak/themes`

4. theme폴더에서 새로운 나만의 폴더 생성

    제시된 양식따라 폴더 및 테마파일 생성

    (기존 base 혹은 keycloak 테마폴더와 같이 mytheme폴더를 생성하여 하위폴더 login 등 양식(폴더위치)를 동일하게 맞춰주기

5. 해당 테마를 적용시키고싶은 Realm Settings >  Themes > 테마선택

    기본은 keycloak, 예시로 보여지는 것은 base, 자신이 폴더를 생성하여 만든 "mytheme"도 카테고리에 보여질 것!

     ![img/Untitled%208.png](img/Untitled%208.png)
     
     
### **Keycloak User 속성 추가**

[https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/](https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/)

 회원가입 시 인자값 ( User attribute 추가)

: 필요한 이유 - 같은 realms  안에 여러 client (Service Provider)가 존재시 각 client별로 로그인 시 필요로하는 metadata값, role 등 다양할 수 있으므로 해당  realm안의 user의 attribute(기본 Property 이외의 속성값)를 매핑시킨다. 

두가지 방법이 있다.

- admin 관리 페이지에서 관리자가 User를 직접 생성

    [https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/](https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/)

    1. 사용자 속성에서 추가 

        `해당 realm > Users > 특정사용자의 Atrributes` - Key 값 쌍 추가 (추가하고싶은 attribute)

         ![img/Untitled%209.png](img/Untitled%209.png)
         
    client 와  mapping 시켜주는 방법 2가지

    - Realm- Configure > Client Scope으로 지정 (여러 client 재사용성)
        1. Client Scope에 추가하여 매핑

             이 시점에서는 keycloak은 속성을 알지만 우리가 사용하고자하는 (Service Provider) client는 속성을 보지 못한다. 따라서 client 용으로 mapping되어야하며 client Scope으로 수행된다.

            `해당 realm > Configure : Client Scopes` - Create

            `추가한 해당 scope > Mappers` - Create 

            :  매핑하고 싶은 이름 / 매퍼유형 : User Attribute / User에서 지정한 key값과 동일하게

        2. 해당 속성을 사용하고자하는 client에게 client scope  접근 권한 부여

            `사용하고자 하는 client > Client Scopes >  Available Client Scopes` -해당scope선택

    - 특정 client >  Mapper로 속성값  mapping

- account service 페이지에서 UI로 회원가입창에서 User가  User를 생성

    [https://www.keycloak.org/docs/latest/server_development/index.html#_custom_user_attributes](https://www.keycloak.org/docs/latest/server_development/index.html#_custom_user_attributes)

    1. 해당 realm > settings > login > register 기능   ON 해주기
    2. 해당 realm > settings > theme  원하는 테마 선택
    3. keycloak container에 접속 
    4. mobile attribute 추가 예시 (링크 참조할것)

        `/opt/jboss/keycloak/themes/<해당theme>/login/register.ftl` 파일 수정

        `/opt/jboss/keycloak/themes/<해당theme>/account/account.ftl` 파일 수정

    5. `http://localhost:8080/auth/realms/<원하는 realm명>/account` 접속

        login 페이지에서 > register > 회원정보입력

---

### Keycloak API 이용

- **Keycloak API**

    Keycloak API 공식페이지 :
     [https://www.keycloak.org/docs-api/12.0/rest-api/#_userrepresentation](https://www.keycloak.org/docs-api/12.0/rest-api/#_uri_scheme)

    parameter 에 대한 부연 설명: 
    [https://www.keycloak.org/docs-api/12.0/rest-api/#_userrepresentation](https://www.keycloak.org/docs-api/12.0/rest-api/#_userrepresentation)

    Test가능한 Postman 공식페이지 : [https://documenter.getpostman.com/view/7294517/SzmfZHnd#e917ce53-69ea-49f3-9a94-4f6c0962c199](https://documenter.getpostman.com/view/7294517/SzmfZHnd#e917ce53-69ea-49f3-9a94-4f6c0962c199)

    - Realm Master에 대한 access_token 조회:

        **POST  Obtain access token for a user**

        url: `http://localhost:8080/auth/realms/master/protocol/openid-connect/token`

        **Headers**: Content-Type : application/x-www-form-urlencoded

        **Body :**  

         ![img/Untitled%2010.png](img/Untitled%2010.png)

        ```bash
        curl --location --request POST 'http://127.0.0.1:8180/auth/realms/master/protocol/openid-connect/token' \
        --header 'Content-Type: application/x-www-form-urlencoded' \
        --data-urlencode 'client_id=admin-cli' \
        --data-urlencode 'username=<master realm login 시 이용하는 username>' \
        --data-urlencode 'password=<master realm login 시 이용하는 password>' \
        --data-urlencode 'grant_type=password'
        ```

    - 회원가입 :

        **POST Create user** 

        url : `http://localhost:8080/auth/admin/realms/<회원가입하고싶은해당realm명>/users`

        **Authorization - Bearer Token** : <조회한 master의  AccessToken>

        **Headers**: Content-Type : application/json

        ```bash
        curl --location --request POST 'http://127.0.0.1:8180/auth/admin/realms/<회원가입하고싶은해당realm명>/users' \
        --header 'Content-Type: application/json' \
        --data-raw '{
                "createdTimestamp": 1588880747548,
                "username": "Strange",
                "enabled": true,
                "totp": false,
                "emailVerified": true,
                "firstName": "Stephen",
                "lastName": "Strange",
                "email": "drstranger@marvel.com",
                "disableableCredentialTypes": [],
                "requiredActions": [],
                "notBefore": 0,
                "access": {
                    "manageGroupMembership": true,
                    "view": true,
                    "mapRoles": true,
                    "impersonate": true,
                    "manage": true
                },
        			  "attributes": {"mobile":"123"},
                "realmRoles": [	"mb-user" ]
            }'
        ```

    - 특정 Realm의 User 조회:

        **GET Get users** 

        url : `http://localhost:8080/auth/admin/realms/<조회하고 싶은 realm명>/users`

        **Authorization - Bearer Token** : <조회한 master의  AccessToken>

        **Headers**: Content-Type : application/json

        (선택사항 : 특정 사용자 조회할때 ) **Params** : Key에 {

        briefRepresentation / email / first / firstName / lastName / max / search / username} 로 검색가능

        ```bash
        curl --location --request GET 'http://127.0.0.1:8180/auth/admin/realms/<조회하고 싶은 realm명>/users'
        ```

    - 특정 User Logout:

        **POST Logout user**

        url: `http://localhost:8080/auth/admin/realms/dudaji-200/users/<조회한 user 결과값 중 "id"값>/logout`

        **Authorization - Bearer Token** : <조회한 master의  AccessToken>

        **Headers**: Content-Type : application/json

        ```bash
        curl --location --request POST 'http://127.0.0.1:8180/auth/admin/realms/<realm명>/users/<조회한 user 결과값 중 "id"값>/logout'
        ```
