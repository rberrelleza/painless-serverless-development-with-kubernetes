# painless-serverless-development-with-kubernetes
Painless Serverless Development With Kubernetes

# Prerequisites
- A [running instance](https://docs.openfaas.com/deployment/kubernetes/) of OpenFaaS
- Docker

# Steps
1. Install faas cli

    ```
    curl -sSL https://cli.openfaas.com | sudo -E sh
    ```

1. Login to your gateway

    ```
    echo $PASSWORD | faas login -g $GATEWAY_URL --password-stdin
    ```
    
    ```
    Calling the OpenFaaS server to validate the credentials...
    credentials saved for admin https://openfaas-ingress-openfaas-rberrelleza.cloud.okteto.net
    ```

1. Pull the required templates
    ```
    faas template pull https://github.com/openfaas-incubator/golang-http-template
    ```

    ```
    Fetch templates from repository: https://github.com/openfaas-incubator/golang-http-template at master
    2019/11/12 16:33:35 Attempting to expand templates from https://github.com/openfaas-incubator/golang-http-template
    2019/11/12 16:33:36 Fetched 4 template(s) : [golang-http golang-http-armhf golang-middleware golang-middleware-armhf] from https://github.com/openfaas-incubator/golang-http-template
    ```
1. Build and launch the function

    ```
    faas up -f hello.yml
    ```

    ```
    [0] > Building hello.
    ...
    ...
    ...
    Deploying: hello.

    Deployed. 202 Accepted.
    URL: https://openfaas-ingress-openfaas-rberrelleza.cloud.okteto.net/function/hello
    ```

1. Call the hello function

    ```
    faas invoke hello -f hello.yml
    ```

    ```
    Reading from STDIN - hit (Control + D) to stop.
    hello
    Hello world, input was: hello
    ```


1. Launch your dev environment

    ```
    cd hello
    okteto up
    ```

1. Change the return message in `hello/handler.go`
    
    ```
    sed -i '' 's/world/Serverless Summit/' function/handler.go
    ```

1. Start the function in your dev environment

    ```
    fwatchdog
    ```
    
    ```
    Forking - go [run main.go]
    2019/11/13 00:31:00 Started logging stderr from function.
    2019/11/13 00:31:00 Started logging stdout from function.
    2019/11/13 00:31:00 OperationalMode: http
    2019/11/13 00:31:00 Timeouts: read: 10s, write: 10s hard: 10s.
    2019/11/13 00:31:00 Listening on port: 8080
    2019/11/13 00:31:00 Writing lock-file to: /tmp/.lock
    2019/11/13 00:31:00 Metrics listening on port: 8081
    ```

1. Call the hello function again

    ```
    faas invoke hello -f hello.yml
    ```

    ```
    Reading from STDIN - hit (Control + D) to stop.
    hello
    Hello Serverless Summit, input was: hello
    ```

