# Gitea - solution

## Network

Create the bridge network `gitea`:

```bash
docker network create -d bridge gitea
```

## Database

To run the database container you need to first create the data directory and then start it:

```bash
mkdir ./mysql_data
docker run -d \
  --name mysql \
  --network gitea \
  --restart unless-stopped \
  -v <Path to current directory>/mysql_data:/var/lib/mysql \
  -e MYSQL_ROOT_PASSWORD=gitea \
  -e MYSQL_DATABASE=gitea \
  -e MYSQL_USER=gitea \
  -e MYSQL_PASSWORD=gitea \
  mysql:8
```

The options in detail:

- `-d`: Run in detached mode (you could close your shell and the db should continue to run)
- `--name mysql`: This is required so that we can latter connect to the DB using this as hostname
- `--network gitea`: Put the container into a separate network (recommended to isolate applications, required to reference other containers by hostname)
- `--restart unless-stopped`: Container will be restared it if fails, only if you explicitly stop it, it doesn't restart
- `-v <Path to current directory>/mysql_data:/var/lib/mysql`: Mounts the directory `mysql_data` in the specified path into the container at `/var/lib/mysql`
- `mysql:8`: We are using the official mysql image from docker.io

To inspect the container use  `ps` and `logs`:

```bash
$ docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS              PORTS                 NAMES
10569cc35821   mysql:8   "docker-entrypoint.s…"   About a minute ago   Up About a minute   3306/tcp, 33060/tcp   mysql
$ docker logs mysql
2023-01-24T10:14:40.433068Z 0 [Warning] [MY-011810] [Server] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
2023-01-24T10:14:40.450263Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /var/run/mysqld/mysqlx.sock
2023-01-24T10:14:40.450294Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.0.32'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.
```

You should see that there is a line saying "ready for connections", then you know that everything worked.

## Gitea

The gitea server can be started like that:

```bash
mkdir ./gitea_data
docker run -d \
  --name gitea \
  --restart unless-stopped \
  --network gitea \
  -p 3000:3000 \
  -p 222:22 \
  -v "<Path to the current directory>"/gitea_data:/data \
  -v /etc/localtime:/etc/localtime:ro \
  -v /etc/timezone:/etc/timezone:ro \
  -e "USER_UID=1000" \
  -e "USER_GID=1000" \
  -e "GITEA__database__DB_TYPE=mysql" \
  -e "GITEA__database__HOST=mysql:3306" \
  -e "GITEA__database__NAME=gitea" \
  -e "GITEA__database__USER=gitea" \
  -e "GITEA__database__PASSWD=gitea" \
  gitea/gitea:1
```

The options in detail:

- `-d`: Run in detached mode (you could close your shell and the db should continue to run)
- `--name gitea`: Makes managing the container easier
- `--network gitea`: Put the container into a separate network (recommended to isolate applications, required to reference other containers by hostname)
- `--restart unless-stopped`: Container will be restared it if fails, only if you explicitly stop it, it doesn't restart
- `-v <Path to current directory>/gitea_data:/var/lib/mysql`: Mounts the directory `gitea_data` in the specified path into the container at `/var/lib/mysql`
- `-v /etc/localtime:/etc/localtime:ro`: Mounts the file `/etc/localtime` into the container at the same path but read-only (the application requires this to set the time & zone accordingly)
- `gitea/gitea:1`: We are using the gitea image from docker.io in version 1 (omits patch and feature version)

## Bonus

Stopped container: We don't specify the `--rm` option, so if we stop the containers we can simply start them again and everyting was as before.

Removed container: If we remove the container, we can start new ones with the same command and get our installed gitea back, since both the DB and gitea store their data outside the container in the mounted directory (very important concept!)
