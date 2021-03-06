# ๐ SSO ๐

**๐   ๋ชฉํ - Gitlab์ Keycloak SAML๋ฐฉ์์ ์ด์ฉํ SSO ํ๊ฒฝ ๊ตฌ์ถ**

  - ๊ตฌ์ถ ํ๊ฒฝ 

    docker-desktop (mac) & kubernetes

    **local**์์ ์๋
    
    [1. Keycloak pod ๋์ฐ๊ธฐ](#keycloak-์ค์น)

    [2. gitlab pod ๋์ฐ๊ธฐ](#gitlab-์ค์น)

    [3. Keycloak -gitlab SSO ํ๊ฒฝ๊ตฌ์ถ](#keycloak---gitlab-์ฐ๋)
     (SAML ๋ฐฉ์์ด์ฉ)
     
    [4. Keycloak Login, Register Customizing](#keycloak-login,-register-customizing)
    
    [5. Keycloak API ์ด์ฉ](#keycloak-API-์ด์ฉ)
---

## Keycloak ์ค์น

local์์ ์๋ํ๋ฏ๋ก ingress ์์ด 

    `kubectl create -f https://raw.githubusercontent.com/keycloak/keycloak-quickstarts/latest/kubernetes-examples/keycloak.yaml`

  ๋๊ฐ์ง ๋ฐฉ๋ฒ ์ค ์ ํ

  1. `kubectl edit svc keycloak` ํด์ loadbalancer ๋ฅผ NodePort๋ก ๋ฐ๊ฟ์ฃผ๋๊ฐ

      Keycloak ํ์ด์ง ์ฃผ์ : `http://localhost:<nodeport>`

  2. kecloack **์ค์น ์ด์ **์ (keycloak์ ๋จผ์ ์ค์นํ๋ฉด pending์ํ๋ก ์ง์๋จ)

      nginx ์ปจํ์ด๋ nodeport๋ก expose

      --> ์ ์ ๋ก  loadbalancer ์ External_IP๊ฐ  localhost๋ก ์ค์ ๋จ

      ~~์ด๋ป๊ฒ ์ด๋ ๊ฒ ์๋๋์ง๋ ๋ชจ๋ฅด๊ฒ ์!! ์๋ local์์๋ loadbalancer์คํ X~~

      Keycloak ํ์ด์ง ์ฃผ์ : http://localhost:8080

---
## Gitlab ์ค์น

๋๊ฐ์ง ๋ฐฉ๋ฒ ์ค ์ ํ (์ฌ๊ธฐ์๋ 2๋ฒ๋ฐฉ๋ฒ ์ด์ฉ)

1.  helm ์ฌ์ฉ

    ์ค์น :[https://docs.gitlab.com/charts/installation/](https://docs.gitlab.com/charts/installation/)

2. image ์ด์ฉํด์ deployment > service : NodePort ๋ก ์ค์ 
    - command ์ด์ฉํ๋ ๋ฐฉ๋ฒ

        `kubectl create deployment gitlab --image=gitlab/gitlab-ce:latest --port=80`

        `kubectl expose deployment gitlab --type=NodePort --port=80 --target-port=80 --name=gitlab-service`

    - yamlํ์ผ์ ์ด์ฉํ๋ ๋ฐฉ๋ฒ (command ๋ฐฉ๋ฒ์ nodeport๋ฅผ ์ง์ ํด์ค์ ์๋ค!)
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

        gitlab์ ์คํ๋ ์ดํ 5๋ถ์ ๋ ์ดํ์ ์ ์๊ฐ๋ฅ

         ๋ฐ๋ก ์งํ์  ~~curl: (52) Empty reply from server์ด๊ฑฐ๋  502 whoops something went wrong on our end ์๋ฌ๋ธ~~

> gitlab์ค์น ํ ์ ์๋์๊ฐ๋ฉด โ> kubernetes  ๋ฌธ์ ์ 
> .  >> docker-desktop์ setting โ resource > memory๋ฅผ ์ฌ๋ ค์ฃผ์ด๋ผ!

---
## Keycloak - Gitlab ์ฐ๋

Keycloak์ 2๊ฐ์ง SAML, OpenId-connect ๋ฐฉ์์ด ์๋ค.

Gitlab์ ์ฌ๋ฌ  ์ข๋ฅ OmniAuth provider์ ํตํฉํ๋ค.  keycloak์ SAML, openid-connect ๋ฐฉ์ ๋ฑ์ OmniAuth provider(ํ์ค identity provider protocol)์ ์ ๊ณตํ๋ค.

์ฌ๊ธฐ์์๋ SAML๋ฐฉ์์ ์ด์ฉํด์ SSO๋ฅผ ๊ตฌํํด๋ณด๊ฒ ๋ค.

### Keycloak setting

[http://localhost:8080/auth](http://localhost:8080/auth) ๋ก
์ฒ์ ์ ๊ทผ์: [master realm = admin]๊ณ์ ์ผ๋ก ์ ๊ทผ (`์์ด๋ admin/ ๋น๋ฒ admin` : keycloak์ค์น์ yamlํ์ผ์์ ์ด๊ธฐ ์์ด๋,๋น๋ฒ ์ง์ ํด์ค)

1. realm ์์ฑ

    : ์ด๊ธฐ์๋ Master Realm๋ง ์กด์ฌ

     but, ์ฐ๋ฆฌ๋ Master๊ฐ ์๋ MSA์์ ์ฌ์ฉํ  Realm์ด ํ์ํ๊ธฐ๋๋ฌธ์ ์์ฑ

    ![img/Untitled.png](img/Untitled.png)

    "Add realm"๋ก ์์ฑ

    Realm Settings ์์ 
    { ์ผ๋ฐ์ค์ , ๋ก๊ทธ์ธ ํ๋ฉด ์ค์ , ํค ์ํธํ ์ค์ , ์ด๋ฉ์ผ ์๋น์ค ์ค์ , ๋ก๊ทธ์ธ/ํ์๊ฐ์ ๋ฑ์ ํ๋ง ์ค์ , ์บ์ ์ค์ , ํ ํฐ ์ค์ } ๋ฑ โโ > ํ์์ ๊ณต์๋ฌธ์ ์ฐธ๊ณ ํ ๊ฒ!
    ~~ํ์๊ฐ์ ๋ฐ ๋ก๊ทธ์ธ , user field  ์ถ๊ฐ๋ ๋ค์์~~

2. user ์์ฑ

    realms ์์์ ๊ด๋ฆฌ์๋ฅผ ํตํด ์ ์ ๋ฅผ ์์ฑํ  ์ ์๋ค.

    <์ฌ์ฉํ๊ณ ์ํ๋ realm>ํ๊ฒฝ์์ ํ์ธ -> Manage -> Users ํด๋ฆญ -> Add User ํด๋ฆญํ๋ค

    ![img/Untitled%201.png](img/Untitled%201.png)

    ์ดํ ์์ฑ๋๋ฉด "credentials"ํญ์์ ์ ์ ์๊ฒ ์์๋น๋ฐ๋ฒํธ ๋ฐ๊ธํด์ค๋ค (Set Password)ํด์ค

3. client ์์ฑ
    gitlab ๊ณผ SAML ๋ก SSO ๋ฅผ ํ ๊บผ๋๊น ์ค์ ํด์ค์ผํจ
    
    - client ์์ฑ ์ "SAML"  Protocol๋ก ์์ฑ >> ์ดํ ์๋์ ๊ฐ์ด ์ค์ 

      ![img/Untitled%202.png](img/Untitled%202.png)

      ![img/Untitled%203.png](img/Untitled%203.png)

      ![img/Untitled%204.png](img/Untitled%204.png)

      IDP initiated SSO URL Name = client๋ช๊ณผ ๊ฐ๊ฒ ํด์ฃผ๊ธฐ

      ํด๋น ์ต์ ๋ณ ์๋ฏธ : [https://www.keycloak.org/docs/latest/server_admin/#saml-clients](https://www.keycloak.org/docs/latest/server_admin/#saml-clients)


    - ํด๋น client์ Mapper ์ค์ ํด์ฃผ๊ธฐ

        MapperType > name,email,first_name,last_name๋ง ๊ธฐ๋ณธ User property

        Property > User์์ฑ์ ๊ฐ ์์ฑ๋ช์ ๋งค์นญ

        SAML Attribute Name > gitlab์ผ๋ก ์ ์ก๋๋ data๋ช์นญ

        ![img/Untitled%205.png](img/Untitled%205.png)
        
         โ> ์คํ ์์ saml attribute name : last_name์ผ๋ก ํ ๊ฒ!

        keycloak์ data๋ฅผ gitlab(SP: Service Provider)์ด ํ์๋ก ํ๋ ๊ฐ์ผ๋ก Mappingํ์ฌ ์ ํํ ์ ์กํ  ์์๋ค. (metadata, role ๋ฑ)

- **(๋ฒ์ธ) keycloak accountํ์ด์ง์์ ๋ก๊ทธ์ธ ๋ฐ ๊ณ์  ํ์ธ๋ฐฉ๋ฒ**

    `http://localhost:8080/auth/realms/<์ฌ์ฉํ๊ณ ์ํ๋ realm>/account` 

    (๊ด๋ฆฌ์ํ์ด์ง)์์ฑํ realm์์ ์์ฑํ user์ ๋ณด๋ก ๋ก๊ทธ์ธํด์ ๊ณ์  ํ์ธ๊ฐ๋ฅ

      
### Gitlab setting

Pod๋ก ์ ์ํ์ฌ KeyCloak์ ์ค์ ๊ฐ์ ๋ฐ๊ฟ์ค

- ๊ณต์ํ์ด์ง [https://docs.gitlab.com/ee/integration/omniauth.html](https://docs.gitlab.com/ee/integration/omniauth.html)
                  [https://docs.gitlab.com/ee/integration/saml.html](https://docs.gitlab.com/ee/integration/saml.html)
- gitlab์ ์ค์ ํ์ผ์ ์ด์ด์ ํธ์ง

    `sudo vim /etc/gitlab/gitlab.rb`

    OmniAuth ๊ณต๊ธ์ ์ค์ 

    ๊ณต์ํ์ด์ง [https://docs.gitlab.com/ee/administration/auth/oidc.html](https://docs.gitlab.com/ee/administration/auth/oidc.html)

    - **์ธ์ฆ์ ์ ๋ณด**

        Realm Settings > Keys> Public Keys์  Certificate์ ๋ณด (ํด๋ฆญํ์ฌ ๊ฐ ๋ณต์ฌํ๊ธฐ)

        ![img/Untitled%206.png](img/Untitled%206.png)

    -  **OmniAuth ๊ณต๊ธ์ ์ค "SAML" ์ฌ์ฉ*

         [https://www.debyum.com/configure-gitlab-with-keycloak/](https://www.debyum.com/configure-gitlab-with-keycloak/)

        [https://edenmal.moe/post/2018/GitLab-Keycloak-SAML-2-0-OmniAuth-Provider](https://edenmal.moe/post/2018/GitLab-Keycloak-SAML-2-0-OmniAuth-Provider/?__cf_chl_captcha_tk__=abd99e96a0b76ba86f74686fd3ca96783d87b8d8-1618449110-0-Ac7yEOdq9cMmxXtZ05mffkm-WEt8xo_NfputD_f27Rqzh4jehkrLXH_jkbLwf7bbJrjdOEY8aJpa2IGwBfeWITecHQiJBG2N1taIvwrSsk9uvPv4EZP76YucqgQyQq-6mXR5x8hp2UG-ODDV-lT26fwpCBFqk3robXjakxTOShSjjSJOLXczyuc5KjOIJfQrCtrc4tsYquXSoUts953hqNyBPC1qObV5emUjzKYG8KULaYfpkYdji8vtX1cZNCPDpd7eKKjJRey6V4Ce8icUqaDf3bVIdMGmbHMBGTcuwRBFJZW5VI-feIe7m7UvbFmQZ0uV2yVqw9CDLOJAtlxwgVS8ZDht1hHM2Znt7dP8sDCOJ7BxkDYVccsjTK5bsg61w5qFZ8eBr3Xrfpw1KpfMtMXeAgp_VUSKeORiC4d2wwb9d0MXsXIqt1-cYb3xUUgocThJkLnPp3FaHktPV_Z-OMUye9UQmWI3_h-fAHzIbKR5lOCg0W6bHYhxfKv0dfCkwc0Lq46jSTAYbB2yxW5mXufairq7ABdKz29Cg1hH_giXL5n6dUGrBgTUgD7l6NqLFkWc9MBjOCoTfmZwt1VgExMDZEiYPsbObcCXgCJpMi-Q2A20D7byhu47SHNg5rbwlg8aIq5KD7nY_MID1qc2aCT7xKa4hON1vgf_86uG-nCW)

        [https://gist.github.com/int128/ab5839d5f59829840a0204f0c8cd8a8b](https://gist.github.com/int128/ab5839d5f59829840a0204f0c8cd8a8b)

        ```ruby
        gitlab_rails['omniauth_enabled'] = true

        #์ฌ์ฉ์๊ฐ ๋จผ์  ๊ณ์ ์ ์๋์ผ๋ก ๋ง๋ค์ง ์๊ณ ๋ SAML์ ์ฌ์ฉํ์ฌ ๊ฐ์
        gitlab_rails['omniauth_allow_single_sign_on'] = ['saml']
        # gitlab_rails['omniauth_auto_sign_in_with_provider'] = 'saml'
        gitlab_rails['omniauth_block_auto_created_users'] = false

        # gitlab_rails['omniauth_auto_link_ldap_user'] = false

        #์ด๋ฉ์ผ ์ฃผ์๊ฐ ์ผ์นํ๋ ๊ฒฝ์ฐ SAML ์ฌ์ฉ์๋ฅผ ๊ธฐ์กด GitLab ์ฌ์ฉ์์ ์๋์ผ๋ก ์ฐ๊ฒฐ
        gitlab_rails['omniauth_auto_link_saml_user'] = true 

        gitlab_rails['omniauth_providers'] = [
            {
              name: 'saml',
              args: {
                       assertion_consumer_service_url: 'http://localhost:30010/users/auth/saml/callback',
                       idp_cert:  "-----BEGIN CERTIFICATE-----\n<...์ธ์ฆ์์ ๋ณด ๋ณต์ฌํด์ ๋ถ์ฌ๋ฃ๊ธฐ....>\n-----END CERTIFICATE-----\n",
                       idp_sso_target_url: 'http://localhost:8080/auth/realms/dudaji-200/protocol/saml/clients/gitlab',
                       issuer: 'gitlab',
                       name_identifier_format: 'urn:oasis:names:tc:SAML:2.0:nameid-format:persistent'
                     },
              label: 'Keycloak Login'
            }
        ]
        ```

    ์ค์ ํ์ผ์ ๋ณ๊ฒฝ ํ ๋ฌด์กฐ๊ฑด ์ค์ ๋ณ๊ฒฝ ๋ช๋ น์ด๋ฅผ ํด์ค์ผํจ `gitlab-ctl reconfigure`

    - log ํ์ธ 
    
        `cat /var/log/gitlab/gitlab-rails/production.log`

***

## Keycloak Login, Register Customizing

### **Keycloak ํ์๊ฐ์ ํ์ด์ง ์ค์ **

ํ์๊ฐ์ ์ฐฝ์ ๋์ฐ๊ณ  ์ถ์ realm ์ ํ  Realm Settings >  Login

 ![img/Untitled%207.png](img/Untitled%207.png)
 
**User registration - ON** : ํ์๊ฐ์ ํ์ด์ง ์กด์ฌํ๊ฒ

(์์์์ ๊ธฐ๋ณธ์ค์ ์์๋ off ๋์ด์์)



### **Keycloak ๋ก๊ทธ์ธ ํ์ด์ง ๊พธ๋ฏธ๊ธฐ**

[https://www.keycloak.org/docs/latest/server_development/#scripts](https://www.keycloak.org/docs/latest/server_development/#scripts)

[https://www.baeldung.com/spring-keycloak-custom-themes](https://www.baeldung.com/spring-keycloak-custom-themes)

1. kubernetes ์์ ์๋ Pod ( Keycloak container)๋ก ์ง์

     ****โผ๏ธ keycloak container ๋ด๋ถ์์๋ vim์ด๋ nano๊ฐ ์๋ค**

          **โ  volume ๋ง์ดํธ๋ฅผ ์ด์ฉํด ํ์ผ์ ์์ฑํ๊ฑฐ๋ ์์ ํด๋ผ!**

     โ ๊ธฐ์กด ํ๋ง์ธ base, keycloak, keycloak2.0 ํด๋ ์์์ ์์ ํ๋ ๊ฒ ๋ณด๋ค 

    ํด๋น ํด๋๋ฅผ ๋ณต์ฌํ์ฌ ์๋ก์ด ํด๋(์๋ก์ด ํ๋ง)๋ฅผ ์์ฑํ์ฌ ์์  ํน์ ์์ฑํ๋๊ฒ์ด ์ข๋ค!

2. `/opt/jboss/keycloak/standalone/configuration/standalone.xml` ํด๋น ํ์ผ์ ๊ฐ์

    ```xml
    <theme>
    	<staticMaxAge>-1</staticMaxAge>
    	<cacheThemes>false</cacheThemes>
    	<cacheTemplates>false</cacheTemplates>
    	...
    </theme>
    ```

    ๋ค์๊ณผ ๊ฐ์ด ๋ฐ๊ฟ๋ผ โ ๊ฐ๋ฐ๋จ๊ณ์์ ์๋ฒ๋ฅผ ๋ค์ ์์ํ์ง ์๊ณ ๋ ๋ณ๊ฒฝ์ฌํญ์ ํจ๊ณผ๋ฅผ ์ฆ์ ํ์ธ

3. keycloak์ theme ํด๋๋ก ์ง์

    #. Keycloak์ ๊ธฐ๋ณธํ์ด์ง `$JBOSS_HOME` โ /opt/jboss/keycloak

    `cd /opt/jboss/keycloak/themes`

4. themeํด๋์์ ์๋ก์ด ๋๋ง์ ํด๋ ์์ฑ

    ์ ์๋ ์์๋ฐ๋ผ ํด๋ ๋ฐ ํ๋งํ์ผ ์์ฑ

    (๊ธฐ์กด base ํน์ keycloak ํ๋งํด๋์ ๊ฐ์ด mythemeํด๋๋ฅผ ์์ฑํ์ฌ ํ์ํด๋ login ๋ฑ ์์(ํด๋์์น)๋ฅผ ๋์ผํ๊ฒ ๋ง์ถฐ์ฃผ๊ธฐ

5. ํด๋น ํ๋ง๋ฅผ ์ ์ฉ์ํค๊ณ ์ถ์ Realm Settings >  Themes > ํ๋ง์ ํ

    ๊ธฐ๋ณธ์ keycloak, ์์๋ก ๋ณด์ฌ์ง๋ ๊ฒ์ base, ์์ ์ด ํด๋๋ฅผ ์์ฑํ์ฌ ๋ง๋  "mytheme"๋ ์นดํ๊ณ ๋ฆฌ์ ๋ณด์ฌ์ง ๊ฒ!

     ![img/Untitled%208.png](img/Untitled%208.png)
     
     
     
     
### **Keycloak User ์์ฑ ์ถ๊ฐ**

[https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/](https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/)

 ํ์๊ฐ์ ์ ์ธ์๊ฐ ( User attribute ์ถ๊ฐ)

: ํ์ํ ์ด์  - ๊ฐ์ realms  ์์ ์ฌ๋ฌ client (Service Provider)๊ฐ ์กด์ฌ์ ๊ฐ client๋ณ๋ก ๋ก๊ทธ์ธ ์ ํ์๋กํ๋ metadata๊ฐ, role ๋ฑ ๋ค์ํ  ์ ์์ผ๋ฏ๋ก ํด๋น  realm์์ user์ attribute(๊ธฐ๋ณธ Property ์ด์ธ์ ์์ฑ๊ฐ)๋ฅผ ๋งคํ์ํจ๋ค. 

๋๊ฐ์ง ๋ฐฉ๋ฒ์ด ์๋ค.

- admin ๊ด๋ฆฌ ํ์ด์ง์์ ๊ด๋ฆฌ์๊ฐ User๋ฅผ ์ง์  ์์ฑ

    [https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/](https://fixes.co.za/keycloak/adding-attributes-to-a-user-in-keycloak/)

    1. ์ฌ์ฉ์ ์์ฑ์์ ์ถ๊ฐ 

        `ํด๋น realm > Users > ํน์ ์ฌ์ฉ์์ Atrributes` - Key ๊ฐ ์ ์ถ๊ฐ (์ถ๊ฐํ๊ณ ์ถ์ attribute)

         ![img/Untitled%209.png](img/Untitled%209.png)
         
    client ์  mapping ์์ผ์ฃผ๋ ๋ฐฉ๋ฒ 2๊ฐ์ง

    - Realm- Configure > Client Scope์ผ๋ก ์ง์  (์ฌ๋ฌ client ์ฌ์ฌ์ฉ์ฑ)
        1. Client Scope์ ์ถ๊ฐํ์ฌ ๋งคํ

             ์ด ์์ ์์๋ keycloak์ ์์ฑ์ ์์ง๋ง ์ฐ๋ฆฌ๊ฐ ์ฌ์ฉํ๊ณ ์ํ๋ (Service Provider) client๋ ์์ฑ์ ๋ณด์ง ๋ชปํ๋ค. ๋ฐ๋ผ์ client ์ฉ์ผ๋ก mapping๋์ด์ผํ๋ฉฐ client Scope์ผ๋ก ์ํ๋๋ค.

            `ํด๋น realm > Configure : Client Scopes` - Create

            `์ถ๊ฐํ ํด๋น scope > Mappers` - Create 

            :  ๋งคํํ๊ณ  ์ถ์ ์ด๋ฆ / ๋งคํผ์ ํ : User Attribute / User์์ ์ง์ ํ key๊ฐ๊ณผ ๋์ผํ๊ฒ

        2. ํด๋น ์์ฑ์ ์ฌ์ฉํ๊ณ ์ํ๋ client์๊ฒ client scope  ์ ๊ทผ ๊ถํ ๋ถ์ฌ

            `์ฌ์ฉํ๊ณ ์ ํ๋ client > Client Scopes > ย Available Client Scopes` -ํด๋นscope์ ํ

    - ํน์  client >  Mapper๋ก ์์ฑ๊ฐ  mapping

- account service ํ์ด์ง์์ UI๋ก ํ์๊ฐ์์ฐฝ์์ User๊ฐ  User๋ฅผ ์์ฑ

    [https://www.keycloak.org/docs/latest/server_development/index.html#_custom_user_attributes](https://www.keycloak.org/docs/latest/server_development/index.html#_custom_user_attributes)

    1. ํด๋น realm > settings > login > register ๊ธฐ๋ฅ   ON ํด์ฃผ๊ธฐ
    2. ํด๋น realm > settings > theme  ์ํ๋ ํ๋ง ์ ํ
    3. keycloak container์ ์ ์ 
    4. mobile attribute ์ถ๊ฐ ์์ (๋งํฌ ์ฐธ์กฐํ ๊ฒ)

        `/opt/jboss/keycloak/themes/<ํด๋นtheme>/login/register.ftl` ํ์ผ ์์ 

        `/opt/jboss/keycloak/themes/<ํด๋นtheme>/account/account.ftl` ํ์ผ ์์ 

    5. `http://localhost:8080/auth/realms/<์ํ๋ realm๋ช>/account` ์ ์

        login ํ์ด์ง์์ > register > ํ์์ ๋ณด์๋ ฅ

---

### Keycloak API ์ด์ฉ

- **Keycloak API**

    Keycloak API ๊ณต์ํ์ด์ง :
     [https://www.keycloak.org/docs-api/12.0/rest-api/#_userrepresentation](https://www.keycloak.org/docs-api/12.0/rest-api/#_uri_scheme)

    parameter ์ ๋ํ ๋ถ์ฐ ์ค๋ช: 
    [https://www.keycloak.org/docs-api/12.0/rest-api/#_userrepresentation](https://www.keycloak.org/docs-api/12.0/rest-api/#_userrepresentation)

    Test๊ฐ๋ฅํ Postman ๊ณต์ํ์ด์ง : [https://documenter.getpostman.com/view/7294517/SzmfZHnd#e917ce53-69ea-49f3-9a94-4f6c0962c199](https://documenter.getpostman.com/view/7294517/SzmfZHnd#e917ce53-69ea-49f3-9a94-4f6c0962c199)

    - Realm Master์ ๋ํ access_token ์กฐํ:

        **POST  Obtain access token for a user**

        url: `http://localhost:8080/auth/realms/master/protocol/openid-connect/token`

        **Headers**: Content-Type : application/x-www-form-urlencoded

        **Body :**  

         ![img/Untitled%2010.png](img/Untitled%2010.png)

        ```bash
        curl --location --request POST 'http://127.0.0.1:8180/auth/realms/master/protocol/openid-connect/token' \
        --header 'Content-Type: application/x-www-form-urlencoded' \
        --data-urlencode 'client_id=admin-cli' \
        --data-urlencode 'username=<master realm login ์ ์ด์ฉํ๋ username>' \
        --data-urlencode 'password=<master realm login ์ ์ด์ฉํ๋ password>' \
        --data-urlencode 'grant_type=password'
        ```

    - ํ์๊ฐ์ :

        **POST Create user** 

        url : `http://localhost:8080/auth/admin/realms/<ํ์๊ฐ์ํ๊ณ ์ถ์ํด๋นrealm๋ช>/users`

        **Authorization - Bearer Token** : <์กฐํํ master์  AccessToken>

        **Headers**: Content-Type : application/json

        ```bash
        curl --location --request POST 'http://127.0.0.1:8180/auth/admin/realms/<ํ์๊ฐ์ํ๊ณ ์ถ์ํด๋นrealm๋ช>/users' \
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

    - ํน์  Realm์ User ์กฐํ:

        **GET Get users** 

        url : `http://localhost:8080/auth/admin/realms/<์กฐํํ๊ณ  ์ถ์ realm๋ช>/users`

        **Authorization - Bearer Token** : <์กฐํํ master์  AccessToken>

        **Headers**: Content-Type : application/json

        (์ ํ์ฌํญ : ํน์  ์ฌ์ฉ์ ์กฐํํ ๋ ) **Params** : Key์ {

        briefRepresentation / email / first / firstName / lastName / max / search / username} ๋ก ๊ฒ์๊ฐ๋ฅ

        ```bash
        curl --location --request GET 'http://127.0.0.1:8180/auth/admin/realms/<์กฐํํ๊ณ  ์ถ์ realm๋ช>/users'
        ```

    - ํน์  User Logout:

        **POST Logout user**

        url: `http://localhost:8080/auth/admin/realms/dudaji-200/users/<์กฐํํ user ๊ฒฐ๊ณผ๊ฐ ์ค "id"๊ฐ>/logout`

        **Authorization - Bearer Token** : <์กฐํํ master์  AccessToken>

        **Headers**: Content-Type : application/json

        ```bash
        curl --location --request POST 'http://127.0.0.1:8180/auth/admin/realms/<realm๋ช>/users/<์กฐํํ user ๊ฒฐ๊ณผ๊ฐ ์ค "id"๊ฐ>/logout'
        ```
