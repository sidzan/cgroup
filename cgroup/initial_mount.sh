#!/bin/bash
check_root(){
retval=""
id=$(id -u)
if [ $id -ne 0 ]
then
	retval="false"
else 
	retval="true"
fi
echo "$retval"
}

retval=$( check_root )
if [ "$retval" == "true" ]
then
	echo "you are root"
else
	echo "you are not root.Must be root To continue"
fi
