#!/bin/bash

if [ -n "$(docker ps -q -f name=go-server)" ]; then
    echo "go-server容器存在，開始關閉舊容器"
    docker stop go-server
    docker rm go-server
fi

docker build -t demo . --no-cache
docker run -p 8088:8088 --name go-server -d demo