#! /bin/bash
echo "This will initially mount a default folder.Will create a cgroup for cpuset and memory"
a=$(cgcreate)
if [ $? -ne 0 ] 
	then
	echo "first install cgroup-bin"
	
	#apt-get install cgroup-bin -y1	
fi

cgcreate -g cpuset,memory:/Batch22

echo "0"> /sys/fs/cgroup/cpuset/Batch22/cpuset.mems
echo "0"> /sys/fs/cgroup/cpuset/Batch22/cpuset.cpus

echo "260000000" >/sys/fs/cgroup/memory/Batch22/memory.limit_in_bytes
