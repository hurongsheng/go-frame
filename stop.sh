#!/bin/bash

ps -ef|grep http/http|grep -v "grep"|awk '{print $2}'|xargs kill -9
ps -ef|grep grpc/grpc|grep -v "grep"|awk '{print $2}'|xargs kill -9
