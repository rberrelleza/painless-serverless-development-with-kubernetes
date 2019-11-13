# painless-serverless-development-with-kubernetes
Painless Serverless Development With Kubernetes

This demo is part of this talk: https://spsna19.sched.com/event/Wb2t/painless-serverless-function-development-in-kubernetes-ramiro-berrelleza-okteto


# Prerequisites
- A [running instance](https://docs.openfaas.com/deployment/kubernetes/) of OpenFaaS
- [Open FaaS CLI](https://docs.openfaas.com/cli/install/)
- [Okteto CLI](https://github.com/okteto/okteto/blob/master/docs/installation.md)
- Docker running in your local machine
- A DockerHub account


# Develop directly in Kubernetes
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
1. Initialize the function
    
    ```
    faas new hello --lang golang-middleware --handler function --gateway $GATEWAY_URL --prefix $DOCKER_ID
    ```

    ```
    Folder: function created.
      ___                   _____           ____
    / _ \ _ __   ___ _ __ |  ___|_ _  __ _/ ___|
    | | | | '_ \ / _ \ '_ \| |_ / _` |/ _` \___ \
    | |_| | |_) |  __/ | | |  _| (_| | (_| |___) |
    \___/| .__/ \___|_| |_|_|  \__,_|\__,_|____/
          |_|


    Function created in folder: function
    Stack file written: hello.yml
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

1. Call the hello function via the `faas invoke` command. Write `hello` and press (Control + D) to send the request.

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
    cd function
    okteto up
    ```

1. Open `function/handler.go` in your favorite IDE, and change the return message:
    
    ```
    w.Write([]byte(fmt.Sprintf("Hello Serverless Summit, input was: %s", string(input))))
    ```

1. Start the function watchdog in your dev environment

    ```
    cd /home/app/handler
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

1. Restore original configuration
    ```
    exit
    okteto down
    ```

    ```
     ✓  Development environment deactivated
    ```

# Debug a function in Kubernetes

1. Launch your dev environment

    ```
    cd function
    okteto up
    ```

1. Configure the function watchdog to launch the debugger

    ```
    cd /home/app/handler
    export fprocess='dlv debug /home/app/handler --listen 0.0.0.0:2345 --api-version 2 --log --headless'
    fwatchdog
    ```

    ```
    Forking - dlv [debug /home/app/handler --listen 0.0.0.0:2345 --api-version 2 --log --headless]
    2019/11/13 01:55:20 Started logging stderr from function.
    2019/11/13 01:55:20 Started logging stdout from function.
    2019/11/13 01:55:20 OperationalMode: http
    2019/11/13 01:55:20 Timeouts: read: 10s, write: 10s hard: 10s.
    2019/11/13 01:55:20 Listening on port: 8080
    2019/11/13 01:55:20 Metrics listening on port: 8081
    2019/11/13 01:55:20 Writing lock-file to: /tmp/.lock
    ```

1. Open this repo in a local instance of VS Code, and add a breakpoint in `function/handler.go` line 10.

1. Press F5 to start the debugger

1. Call the hello function via the `faas invoke` command. Write `hello` and press (Control + D) to send the request.

    ```
    faas invoke hello -f hello.yml
    ```

    ```
    Reading from STDIN - hit (Control + D) to stop.
    hello
    ```

1. The debugger will stop on the breakpoint. At this point you can control the flow, inspect values and continue the execution.