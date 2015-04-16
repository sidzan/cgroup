
function getcpid() {
    cpids=`pgrep -P $1|xargs`
#    echo "cpids=$cpids"
    for cpid in $cpids;
    do
        echo "$cpid"
        getcpid $cpid
    done
}

getcpid $1
