# Use the alpine base image
#  with the centos base image the container is 350MB big!
#  but now with the alpine base image the container is only 9.1MB small
FROM alpine:latest

# Set the maintainer of the dockerfile
#  on older versions the command "MAINTAINER" is used
#  and this is depricated! The new way to set
#  the maintainer is with the "LABEL" (see below)
LABEL maintainer="nathanael.liechti@post.ch"

# Set the working directory to /var/www/localhost/htdocs/
WORKDIR /var/www/localhost/htdocs/

# Install the apache webserver
RUN apk add --no-cache apache2

# Copy the html file into the container at /var/www/localhost/htdocs/
#  with ah '.' you can add a whole direcotry to the docker container
COPY index.html /var/www/localhost/htdocs/

# Change the permissions of the HTML file (755=rwx,rx,rx)
RUN chmod 755 index.html

# Set the entrypoint where the CMD command start to run 
ENTRYPOINT ["/usr/sbin/httpd"]

# Run the command to start the httpd server as a daemon in the foreground
CMD ["-D", "FOREGROUND"]

# Expose Port 80 for the apache webserver
EXPOSE 80