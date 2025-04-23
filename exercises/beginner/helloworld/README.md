# Hello World

In this exercise you should create a new Dockerfile that starts a bash script and then exits. The script is already provided to you (see the [script.sh](./script.sh))

Here are the different steps:

- use `alpine:latest` as base
- set the workdir to `/`
- install `bash`
- copy the script.sh into the container
- set the permissions on the shell script to 755
- set the CMD directive to run the script at start

Build an image and run a container in foreground and check that "Hello World" get's printed.

## Bonus

Try to find a faster way to print "Hello World" with a container without building your own Dockerfile.

## Solution

See the working [Dockerfile](./Dockerfile) and [here](./solution.md) for the commands to build and run in.
