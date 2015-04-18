#! /bin/bash
cgcreate -g cpuset,memory:/Batch22
echo "0"> /sys/fs/cgroup/cpuset/Batch22/cpuset.mems
echo "0"> /sys/fs/cgroup/cpuset/Batch22/cpuset.cpus
echo "260000000" >/sys/fs/cgroup/memory/Batch22/memory.limit_in_bytes
