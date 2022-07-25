#!/usr/bin/env bash

configpath=`pwd`"/zest.yaml"

cd service/module

cd $1
echo $configpath

buildError=$(go build 2>&1)
if [ !$buildError ];then
	nohup `./$1 -service $1 -sid $2 -configfile $configpath` >/dev/null 2>&1 &
else
	echo $buildError
	exit 1
fi


exit 1