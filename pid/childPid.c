#include <stdio.h>                
#include <signal.h>
#include <unistd.h>
 
int main(int argc, char  *argv[])
{
    pid_t ppid,pid,spid;    
    int io[2];
    pipe(io);
    if((pid=fork())>0){//father
        close(io[1]);
        read(io[0],&spid,sizeof(pid_t));
        ppid    = getpid();
        printf("father:%ld\tme:%ld\tson:%ld\n",ppid,pid,spid);
    }if(pid==0){//me
        close(io[0]);
        pid = getpid();
        ppid
            = getppid();
        if((spid=fork())>0){//me
 
        }else if(spid==0){//son
            spid    = getpid();
            write(io[1],&spid,sizeof(spid));
        }
    }
 
    return 0;
}
