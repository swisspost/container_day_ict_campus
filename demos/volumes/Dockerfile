FROM nginx
LABEL maintainer="janick.kaenzig@post.ch"
# Copy virtual server
COPY default.conf /etc/nginx/conf.d/
# Copy the html file into the container at /usr/share/nginx/html/
COPY html/index.html /usr/share/nginx/html/
EXPOSE 8080
