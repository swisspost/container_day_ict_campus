# Webserver HA - solution

## Step 1

See the three files [index1.html](./index1.html), [index2.html](./index2.html) and [index3.html](./index3.html)

## Step 3

Use the following commands to run the three webservers:

```
docker run -d --name webserver1 --net webserver-ha  webserver:latest

docker run -d --name webserver2 --net webserver-ha  webserver:latest

docker run -d --name webserver3 --net webserver-ha  webserver:latest
```

## Step 4

To run the reverse proxy use this command:

```bash
docker run -d \
  --name caddy \
  --net webserver-ha \
  -p 8080:80 \
  -v <Path to current directory>/Caddyfile:/etc/caddy/Caddyfile \
  caddy:latest
```
