## Chapter 5 homework

#### p1
```Bash
➜  ch5 git:(master) ./p1
hello world (pid:83504)
hello, I am parent of 83505 (pid:83504)
hello, I am child (pid:83505)
```

In this example, the main program (`p1.c`) calls fork(), prints "hello" and is considered completed; the fork also prints hello , separate from the original pid

#### p2
```Bash
➜  ch5 git:(master) ./p2
hello world (pid:83603)
hello, I am child (pid:83604)
hello, I am parent of 83604 (wc:83604) (pid:83603)
```

In `p2.c`, the main process sleeps for one second, allowing the forked child process to print their hello statement before the main process is able to do so. 

#### p3
```Bash
➜  ch5 git:(master) ./p3
hello world (pid:83636)
hello, I am child (pid:83637)
      32     123     966 p3.c
hello, I am parent of 83637 (wc:83637) (pid:83636)
```

The `exec()` system call allows for `p3.c` to change the current process into a different program - a call to `exec()` never returns. `p3.c` creates a child process and then uses the API to run `wc` in the process (printing the results). The parent then remind us that it still exists by printing its own hello statement.

#### p4
```Bash
➜  ch5 git:(master) cat p4.output
      32     109     846 p4.c
```

Line 14 in `p4.c` contains the following, of which separates it from `p3.c`:
```C
open("./p4.output", O_CREAT|O_WRONLY|O_TRUNC, S_IRWXU);
```
Which makes the appropriate system calls to create a file that the child process uses to run the `wc` program (close stdout, create and write to file `p4.output`), while the parent process continues to run.
