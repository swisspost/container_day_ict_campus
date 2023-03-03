# Use the alpine base image
#  with the centos base image the container is 350MB big!
#  but now with the alpine base image the container is only 9.1MB small
FROM alpine:latest

# Set the maintainer of the dockerfile
LABEL maintainer="janick.kaenzig@post.ch"

# Install the nginx webserver
RUN apk add --no-cache nginx

# Copy config file into the container at /etc/nginx/ 
COPY nginx.conf /etc/nginx/

# Copy the html file into the container at /usr/share/nginx/html/
COPY index.html /usr/share/nginx/html/

# Set the entrypoint where the CMD command start to run 
ENTRYPOINT ["nginx"]

# Run the command to start the nginx in the foreground
CMD ["-g", "daemon off;"]

# Expose Port 8080 for the nginx webserver
EXPOSE 8080