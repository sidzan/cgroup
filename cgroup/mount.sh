#! /bin/bash
echo "This will initially mount a default folder.Will create a cgroup for cpuset and memory"
docker run --rm busybox:latest
cgcreate -g cpuset,memory:/Batch22

echo "0"> /sys/fs/cgroup/cpuset/Batch22/cpuset.mems
echo "0"> /sys/fs/cgroup/cpuset/Batch22/cpuset.cpus

