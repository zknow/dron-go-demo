#!/bin/bash

if [ -n "$(docker ps -q -f name=go-server)" ]; then
    echo "go-server容器存在"
    docker stop go-server
    docker rm go-server
else
    docker build -t demo . --no-cache
    docker run -p 8088:8088 --name go-server -d demo
fi