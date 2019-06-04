# Docker Examples

Here is the documentation about the two dockerfiles "webserver" and "helloworld". 
Please follow the description below to build and run the docker image

### hello World Example:
Change to the webserver directory:  
```
cd example-helloworld  
```

Build the new "helloworld" docker image with the tag "latest":  
```
docker build -t "helloworld:latest" .  
```

Run the docker container with "docker run":  
```
docker run helloworld:latest  
```

Stop and delete the running docker container:  
```
docker stop <container_id>  
docker rm <container_id> 
```

  
### webserver Example:
Change to the webserver directory:  
```
cd example-webserver  
```

Build the new "webserver" docker image with the tag "latest":  
```
docker build -t "webserver:latest" . 
```

Run the docker container with "docker run":  
```
docker run -p 8000:80 -d webserver
```

Stop and delete the running docker container:  
```
docker stop <container_id>  
docker rm <container_id>  
```

Start the docker container with the docker-compose file (for automation):  
```
docker-compose -f /docker_data/docker-ict-campus-demo/example-webserver/docker-compose.yml up -d
```

Stop the docker container with the docker-compose file:  
```
docker-compose -f /docker_data/docker-ict-campus-demo/example-webserver/docker-compose.yml down
```