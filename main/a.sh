#! /bin/sh
(( `id -u` == 0 )) || (echo "Must be root to run script"; exit 1; )

if ( `id -u` -eq 0 )
	then
	echo "hello"
else echo "bad karma"
fi


