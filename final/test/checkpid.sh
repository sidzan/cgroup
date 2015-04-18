#! /bin/bash
echo "$(pgrep -P $(cat /var/run/docker.pid))" > /sys/fs/cgroup/cpuset/Batch22/tasks

echo "$(pgrep -P $(cat /var/run/docker.pid))" > /sys/fs/cgroup/memory/Batch22/tasks
