# Webserver HA - solution

## Step 1

`docker network create webserver-ha`

## Step 2

Use the following commands to run the three webservers:

```
docker run -d --name webserver1 --net webserver-ha  webserver:latest

docker run -d --name webserver2 --net webserver-ha  webserver:latest

docker run -d --name webserver3 --net webserver-ha  webserver:latest
```

## Step 3

To run the reverse proxy use this command:

```bash
docker run -d \
  --name caddy \
  --net webserver-ha \
  -p 8080:80 \
  -v <Path to current directory>/Caddyfile:/etc/caddy/Caddyfile \
  caddy:latest
```
