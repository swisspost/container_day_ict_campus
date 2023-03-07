# JSWiki - solution

To start the wikijs container you need the following command:

```bash
docker run -d \
  --name wiki \
  -p 3000:3000 \
  -e "DB_TYPE=sqlite" \
  -e "DB_FILEPATH=/wiki/wiki.sqlite" 
  -v "<Path to the current directory>/wiki.sqlite:/wiki/wiki.sqlite" 
  ghcr.io/requarks/wiki:2
```

The options in detail:

- `-d`: Run in detached mode (you could close your shell and the wiki should continue to run)
- `--name wiki`: With a name it's easier to stop/start the container later
- `-p 3000:3000`: exposes port 3000 on the host and redirects you to the container's port 3000 (where wikijs is listening on by default)
- `-e DB_TYPE=sqlite`: Sets the environment variable inside the container
- `-e DB_FILEPATH=/wiki/wiki.sqlite`: Specifies the location where wikijs should place the database inside the container
- `-v <Path to the current directory>/wiki.sqlite:/wiki/wiki.sqlite`: Mounts the file wiki.sqlite from the current directory into the container at the location wikijs expects it. This works, even when the file does not yet exist.
- `ghcr.io/requarks/wiki:2`: The image that is started

Stop and delete the running container with it´s name:

```bash
docker stop wiki
docker rm wiki
```

## Bonus

The only file that needs to be backed up is the `wiki.sqlite` file you mounted into the container. All data within the container is ephemeral and can be treated like that.

If you run the same command on another docker host, with the `wiki.sqlite` file copied there first, the wiki starts just as you would expect.

## Bonus 2

First create a volume for postgres and the network:

```bash
docker volume create postgres_data
docker network create wikijs_postgres
```

then run a new postgres database:

```bash
docker run -d \
  --name postgres \
  --net wikijs_postgres \
  --restart unless-stopped \
  -e POSTGRES_DB=wiki \
  -e POSTGRES_PASSWORD=wikijsrocks \
  -e POSTGRES_USER=wikijs \
  -v postgres_data:/var/lib/postgresql/data \
  postgres:11-alpine
```

And then run three wiki instances with the same postgres as backend:

```bash
docker run -d \
  -p 3001:3000 \
  --name wiki_postgres1 \
  --net wikijs_postgres \
  --restart unless-stopped \
  -e "DB_TYPE=postgres" \
  -e "DB_HOST=postgres" \
  -e "DB_PORT=5432" \
  -e "DB_USER=wikijs" \
  -e "DB_PASS=wikijsrocks" \
  -e "DB_NAME=wiki" \
  -e "HA_ACTIVE=true" \
  ghcr.io/requarks/wiki:2
docker run -d \
  -p 3002:3000 \
  --name wiki_postgres2 \
  --net wikijs_postgres \
  --restart unless-stopped \
  -e "DB_TYPE=postgres" \
  -e "DB_HOST=postgres" \
  -e "DB_PORT=5432" \
  -e "DB_USER=wikijs" \
  -e "DB_PASS=wikijsrocks" \
  -e "DB_NAME=wiki" \
  -e "HA_ACTIVE=true" \
  ghcr.io/requarks/wiki:2
docker run -d \
  -p 3003:3000 \
  --name wiki_postgres3 \
  --net wikijs_postgres \
  --restart unless-stopped \
  -e "DB_TYPE=postgres" \
  -e "DB_HOST=postgres" \
  -e "DB_PORT=5432" \
  -e "DB_USER=wikijs" \
  -e "DB_PASS=wikijsrocks" \
  -e "DB_NAME=wiki" \
  -e "HA_ACTIVE=true" \
  ghcr.io/requarks/wiki:2
```

## Super Bonus

The Caddyfile would look something like that:

```Caddyfile
:3333 {
  reverse_proxy  {
    to http://wiki_postgres1:3001 http://wiki_postgres2:3002 http://wiki_postgres3:3003 
    lb_policy random
  }
}
```

And the new reverse_proxy would be started like that:

```bash
docker run -d \
  --name wiki_caddy \
  --net wikijs_postgres \
  -p 3333:3333 \
  -v <Path to current directory>/Caddyfile:/etc/caddy/Caddyfile \
  caddy:latest
```
