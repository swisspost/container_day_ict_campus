# hello World Example:

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

