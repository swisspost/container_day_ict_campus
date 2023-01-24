# Hello World - solution

Build the new "helloworld" docker image with the tag "latest":  
```
docker build -t "helloworld:latest" .  
```

Run the docker container with "docker run":  
```
docker run helloworld:latest  
```

## Bonus

`docker run alpine:latest "echo" "Hello World"` ;)
