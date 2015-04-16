#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>
#include <sys/wait.h>
int main ()
{
pid_t child_pid; /* Create a child process. */
child_pid = fork ();
if (child_pid == 0)
{
exit (0); /* This is the child process. Exit immediately. */
}
else
{
sleep(3); /* This is the parent process. Sleep for a minute. */
system("ps –f f ");
wait(&child_pid);
}
return 0;
} 
