#!/usr/bin/env bash


kill -9 `ps -ef |grep "service" | grep $1|grep $2 | grep -v grep | awk '{print $2}'`