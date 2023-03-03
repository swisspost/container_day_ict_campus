# webserver - solution

Build the new "webserver" docker image with the tag "latest":  
```
docker build -t "webserver:latest" . 
```

Run the docker container with "docker run":  
```
docker run -d --name nginx-webserver -p 8080:8080 webserver:latest
```

Find the container's id:

```bash
$ docker ps
CONTAINER ID   IMAGE                                     COMMAND                  CREATED         STATUS                  PORTS                                   NAMES
645255d3d321   ghcr.io/swisspost/demo-webserver:latest   "nginx -g 'daemon of…"   2 seconds ago   Up Less than a second   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   nginx-webserver
```

Stop and delete the running docker container:  
```
docker stop 645255d3d321
docker rm 645255d3d321
```