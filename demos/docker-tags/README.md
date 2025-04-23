# docker-tags

Dead-simple go app running in container image based on scratch. 

The container takes a build argument to compile the app version in, this can be used to demonstrate docker tags:


```console
docker build -t gogogo .
docker run --rm --name gogogo -p 8080:8080 gogogo
docker build -t gogogo:develop --build-arg VERSION=develop .
docker run --rm --name gogogo -p 8080:8080 gogogo:develop
docker build -t gogogo:v0.0.1 --build-arg VERSION=v0.0.1 .
docker run --rm --name gogogo -p 8080:8080 gogogo:v0.0.1
```
