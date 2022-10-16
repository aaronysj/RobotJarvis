#!/bin/bash

go clean
PID=$(ps -ef |grep Jarvis | grep -v grep |awk '{print $2}')
if [ "${PID}" != "" ]; then
    kill -9 ${PID}
    echo "Kill PID ${PID}"
fi
go build -o Jarvis *.go
nohup ./Jarvis >> output.log &
