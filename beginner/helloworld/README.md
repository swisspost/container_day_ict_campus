# Hello World

In this exercise you should create a new dockerfile that prints "Hello World" to stdout when the container is started and then exits.

Here are the different steps:
- use `alpine:latest` as base
- set the workdir to `/`
- install `bash`
- copy the hello_world_bash_script.sh into the container 
- set the permissions on the shell script to 751
- set the CMD directive to run the script at start

Run the container in foreground and check that "Hello World" get's printed.

## Bonus

Try to find a faster way to print "Hello World" to stdout with a container

## Solution

See the working [Dockerfile](./Dockerfile) and [here](./solution.md) for the commands to build and run in.