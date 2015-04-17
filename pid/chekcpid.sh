#! /bin/bash
#process_id=`/bin/ps -fu $USER| grep -i "docker" | grep -v "grep" | awk '{print $2}'`
#process_id=$(cat /var/run/docker.pid)
#echo $process_id

echo "$(pgrep -P $(cat /var/run/docker.pid))" > /sys/fs/cgroup/cpuset/Batch22/tasks

