# Env

This app reads the env vars `ADDR` and `GREETING` and uses them. They have a default defined in the Dockerfile that can be overriden when starting the container:

```console
docker build -t gogogo .
docker run --rm --name gogogo -p 8080:8080 gogogo
docker run --rm --name gogogo -p 8080:8080 -e GREETING=coffee gogogo
```