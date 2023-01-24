# webserver - solution

Build the new "webserver" docker image with the tag "latest":  
```
docker build -t "webserver:latest" . 
```

Run the docker container with "docker run":  
```
docker run -p 8000:80 -d webserver
```

Find the container's id:

```bash
$ docker ps
CONTAINER ID   IMAGE                                     COMMAND                  CREATED         STATUS                  PORTS                                   NAMES
645255d3d321   ghcr.io/swisspost/demo-webserver:latest   "/usr/sbin/httpd -D …"   2 seconds ago   Up Less than a second   0.0.0.0:8080->80/tcp, :::8080->80/tcp   hungry_davinci
```

Stop and delete the running docker container:  
```
docker stop 645255d3d321
docker rm 645255d3d321
```