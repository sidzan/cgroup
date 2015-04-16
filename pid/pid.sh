#!/bin/sh
pid=$1
if [ -z $pid ]
then
    read -p "PID: " pid
fi
ps -p ${pid:-$$} -o ppid=
