# Use the alpine base image
#  with the centos base image the container is 202MB big!
#  but now with the alpine base image the container is only 9.5MB small
FROM alpine:latest

# Set the maintainer of the dockerfile
#  on older versions the command "MAINTAINER" is used
#  and this is depricated! The new way to set
#  the maintainer is with the "LABEL" (see below)
LABEL maintainer="patrick.buehlmann.1@post.ch"

# Set the working directory inside of the container
# The WORKDIR instruction sets the working directory for any 
#  RUN, CMD, ENTRYPOINT, COPY and ADD instructions that follow it in the Dockerfile.
#  You should use WORKDIR instead of proliferating instructions 
#  like RUN cd … && do-something, which are hard to read, troubleshoot, and maintain.
WORKDIR /

# Install bash, because the alpine base image is 2MB small and has no bash installed
RUN apk add --no-cache bash

# Copy the bash script into the container at /
#  with ah '.' you can add a whole direcotry to the docker container
COPY hello_world_bash_script.sh /

# Change the permissions of the script (751=rwx,rx,x)
RUN chmod 751 hello_world_bash_script.sh

# Run hello_world_bash_script.sh when the container launches
CMD "/hello_world_bash_script.sh"

# the same above in an on-liner:
# docker run alpine:latest "echo" "Hello World"
