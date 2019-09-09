# koki-api

API for KOKi App Platform

## Project Setup

### Requirement

1. GoLand IDE latest
1. Golang 1.9.2
1. MySQL 5.7.8+
1. Docker latest

### How To

#### Setup RabbitMQ
1. Run Command

    ```bash
    # RabbitMQ AMQP Port: 5672
    # RabbitMQ Management Port: 15672
    docker run \
        -d \
        --name rabbitmq \
        -p 15672:15672 \
        -p 5672:5672 \
        -e RABBITMQ_DEFAULT_USER=guest \
        -e RABBITMQ_DEFAULT_PASS=guest \
        saggafarsyad/rabbitmq-delayed:3.7-management
    ```

#### Setup Centrifugo

1. Prepare config.json

    ```
    {
        "admin_password": "admin1",
        "admin_secret": "secret",
        "secret": "secret",
        "namespaces": [
            {
                "name": "debug",
                "publish": true,
                "presence": true,
                "join_leave": true,
                "watch": true,
                "anonymous": true,
                "history_size": 10,
                "history_lifetime": 30,
                "recover": true,
                "log_level": "debug"
            },
            {
                "name": "public",
                "publish": true,
                "presence": true,
                "anonymous": true
            },
            {
                "name": "chat",
                "publish": true,
                "presence": true,
                "join_leave": true,
                "history_size": 10,
                "history_lifetime": 30,
                "recover": true
            },
            {
                "name": "track",
                "publish": true,
                "presence": true,
                "join_leave": true,
                "history_size": 10,
                "history_lifetime": 30,
                "recover": true
            }
        ]
    }
    ```

1. Run Command

    ```
    docker run -d \
        --name koki-centrifugo \
        --ulimit nofile=65536:65536 \
        -v $PWD/centrifugo:/centrifugo \
        -p 8002:8000 \
        cr.nbs.co.id/docker/images/centrifugo:1.8.0 centrifugo \
        -c config.json \
        --web
    ```

#### Clone Project

1. Create Go workspace folder in `<YOUR_PROJECT_FOLDER>/src/repo.nusantarabetastudio.com/koki-app`
1. Clone project to `<YOUR_PROJECT_FOLDER>/src/repo.nusantarabetastudio.com/koki-app`
    ```bash
    cd YOUR_PROJECT_FOLDER/src/repo.nusantarabetastudio.com/koki-app
    git clone git@repo.nusantarabetastudio.com:koki-app/koki-api.git
    ```
1. Set `$GOPATH` workspace to `YOUR_PROJECT_FOLDER` path
    - Goland: Go to **Settings > Go > GOPATH > Project GOPATH**, add using **+** button in right side of the window frame.
    - Linux/Mac Terminal:
     ```bash
     cd YOUR_PROJECT_FOLDER
     export GOPATH=$PWD
     ```
1. Open terminal and run command `go get -v ./...`, make sure `$GOPATH` has been set to `YOUR_PROJECT_FOLDER` using command `echo $GOPATH`

#### Run Project

1. Create config file `config.yml` in project source folder/executable folder
1. Create datasource config file `datasources.yml` in project source folder/executable folder
1. Run project with arguments `-c config.yml` and environment variables:
    - `KOKI_NODE_NO=1`

### Deploying Project

1. Run `build.sh` script
1. Copy files in `$GOPATH/bin/koki-api` to server (e.g. `~/koki-api`)
1. Re-configure:
    - `config.yml`:
        - Make sure `auth.email_verification_url` is pointing to address where the server will be deployed.
        - Make sure `auth.email_verification_redirect` is set to the right landing page.
    - `datasources.yml`:
        - Make sure `koki_db` is not pointing to `localhost`.
1. Create `Dockerfile` e.g. in `~/koki-docker`

    ```Dockerfile
    FROM golang:1.9-stretch
    LABEL maintainer="Saggaf Arsyad <saggaf@nusantarabetastudio.com>"
    ENV KOKI_NODE_NO 1
    RUN mkdir -p "/usr/bin/koki"
    WORKDIR /usr/bin/koki
    CMD ["./api", "-c", "config.yml"]
    ```

1. Build Docker image

    ```bash
    cd ~/koki-docker
    docker build -t koki/api:dev .
    ```

1. Build docker container

    ```bash
    cd ~/koki-api
    docker run -d --name=dev_koki_api -v $(pwd):/usr/bin/koki --restart on-failure -p 8081:8081 koki/api:dev
    ```

### File Templates

#### `config.yml`

```yaml
server:
  mode: development
  port: 8080
  base_url: /v1
  default_lat: <DEFAULT_LAT>
  default_lng: <DEFAULT_LNG>

auth:
  dash_api_key: <ADMIN_DASHBOARD_API_KEY>
  jwt_key: <JWT_KEY>
  jwt_lifetime: 2880 # in minutes
  jwt_audience: KOKi Middleware
  jwt_issuer: KOKi Security Authority
  client_secret: <CLIENT_SECRET>
  email_verification_url: <API_BASE_URL>/auth/email/verify?v=%s&t=%d # Must contains exact query parameters
  email_verification_redirect: <REDIRECT_URL>

logger:
  level: debug

base_url:
  avatar: /base/url/to/avatars/
  banner: /base/url/to/banners/
  chef_cover: /base/url/to/chef-covers/
  menu: /base/url/to/menus/
```

#### `datasources.yml`

```yaml
rdbms:
  - name: koki_db
    driver: mysql
    host: <DB_HOST>
    port: <DB_PORT>
    username: <DB_USERNAME>
    password: <DB_PASSWORD>
    database: <DB_NAME>
    max_idle_connection: 10      # optional, set to 10 if unset
    max_open_connection: 10      # optional, set to 10 if unset
    max_connection_lifetime: 1   # seconds, optional, set to 1 second if unset
  - name: koki_cms_db
    driver: mysql
    host: <CMS_DB_HOST>
    port: <CMS_DB_PORT>
    username: <CMS_DB_USERNAME>
    password: <CMS_DB_PASSWORD>
    database: <CMS_DB_NAME>
    max_idle_connection: 10      # optional, set to 10 if unset
    max_open_connection: 10      # optional, set to 10 if unset
    max_connection_lifetime: 1   # seconds, optional, set to 1 second if unset

storage-s3:
  - name: koki_s3
    access_key: <AWS_S3_ACCESS_KEY>
    secret_key: <AWS_S3_SECRET_KEY>
    region: <AWS_S3_REGION>
    bucket_name: <AWS_S3_BUCKET_NAME>
    root_folder: /path/to/root/assets/
```

#### `components.yml`

```yaml
- name: nexmo
  type: nexmo
  configuration:
    base_url_v1: https://rest.nexmo.com
    base_url_v2: https://api.nexmo.com
    api_key: <NEXMO_API_KEY>
    api_secret: <NEXMO_API_SECRET>
    brand: KOKi
    sender_id: KOKi
- name: mailgun
  type: mailgun
  configuration:
    domain: mg.koki.id
    private_validation_key: <MAILGUN_PRIVATE_VALIDATION_KEY>
    public_api_key: <MAILGUN_PRIVATE_API_KEY>
    template_path: /absolute/path/to/template/folder
- name: fcm
  type: push-notifier
  configuration:
    provider: fcm
    server_key: <FCM_SERVER_KEY>
- name: doku
  type: doku-payment-gateway
  configuration:
    mall_id: <DOKU_MALL_ID>
    shared_key: <DOKU_SHARED_KEY>
    timeout: 10000 # in millis
    va_prefix: '12345678' # 8 digit prefix from DOKU
- name: mq
  type: rabbitmq
  configuration:
    username: <RABBITMQ_USERNAME>
    password: <RABBITMQ_PASSWORD>
    host: <RABBITMQ_HOST>
    port: <RABBITMQ_PORT> # e.g. 5672
- name: ws
  type: centrifugo
  configuration:
    private_channel_prefix: $ # private channel prefix usually $
    user: <CENTRIFUGO_CLIENT_USER_ID> #e.g. server
    info: <CENTRIFUGO_CLIENT_USER_INFO> #e.g use empty string '' or json string '{"name" : "Backend Server"}'
    timeout: 10000 # in millis
    secret_key: <CENTRIFUGO_SECRET>
    ws_url: <CENTRIFUGO_WS_URL> # usually ws://localhost:8000/connection/websocket
    api_url: <CENTRIFUGO_API_URL> # usually http://localhost:8000/api/
- name: gmap
  type: googlemap
  configuration:
    api_key: <GOOGLE_MAP_API_KEY>
```

### Assets Folder Structure

```
→ /assets/
   ↳ avatars
   ↳ banners
   ↳ chef-covers
   ↳ menus
```

### Contributors

- Saggaf Arsyad <saggaf@nusantarabetastudio.com>