#! /bin/sh
echo $1

docker inspect --format '{{.State.Pid}}' $1
