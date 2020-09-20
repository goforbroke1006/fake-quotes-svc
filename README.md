# fake-quotes-svc



### How to use

1) Clone repository
2) Enter project directory
3) Run
    ```bash
    # create config file, change settings if you want
    cp config.yaml.dist config.yaml
    
    # build image accordingly your OS and processor architecture 
    make image/local
    
    # run local env
    docker-compose up 
    ```
4) Open http://localhost:18082/
5) Open browser console. Now you see quotes streaming.
