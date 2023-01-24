# Gitea

[Gitea](https://gitea.io/en-us/) is a self-hosted git-server with a cup of tea. It's written in go, super-fast and can work with sqlite for lightweight setups.

In this guide you're going to setup gitea with a separate database container.

For a reference, [here's](https://docs.gitea.io/en-us/install-with-docker/) their official documentation how to install it using docker.

Note: they use docker-compose for the installation, you can just ignore this, we will later discuss about what this is.

## Network

We haven't talked about [docker networks](https://docs.docker.com/network/) so far, but think of them as separate virtual networks, just as you got multiple subnets in your virtualized environment. They are used to separate multiple containerized apps from each other and have the benefit that containers can directly resolve each other by their hostname (that's why it's important to give them a name using --name).

For this app we are going to create our own network using `docker network create gitea`. 

You can find a list of all networks using `docker network ls`.

To start a container in a separate network, run it using the `--net network_name` option.

## Database

First you need a database container to connect to. We are using mysql for this here, but gitea supports [many database engines](https://docs.gitea.io/en-us/database-prep/).

To start the mysql container configure the following options:

- Detached mode
- named `mysql` (so that we can connect to it by NAME:PORT)
- Network: gitea
- using image `mysql:8`
- set [restart-policy](https://docs.docker.com/config/containers/start-containers-automatically/) to `unless-stopped`
- data in `/var/lib/mysql` should be mounted into the container from a local `./mysql_data` directory.
- The following environment variables are required to configure mysql:
  - `MYSQL_ROOT_PASSWORD=gitea`: Password of the user `root` that can configure the entire database server
  - `MYSQL_DATABASE=gitea`: Name of the database the container will create
  - `MYSQL_USER=gitea`: Username for the user that will be able to access the created database
  - `MYSQL_PASSWORD=gitea`: Password for the gitea user

Write your `docker run ...` command and see if the container comes up. Inspect the logs to see if it worked.

Once the container is running, you can proceed.

## Gitea

With the DB up, it's time to spin up the gitea server.

You must configure the following:

- Name: gita
- Image: gitea/gitea:1
- Network: gitea
- Restart Policy unless-stopped
- Port 3000 exposed to the host as 3000
- Port 22 exposed to the host as 222
- A local `./gitea_data` directory mounted in `/data`
- The host-file `/etc/localtime` and `/etc/timezone` mounted inside the container at the same location as read-only
- The following environment variables inside the container:
  - `USER_UID=1000`
  - `USER_GID=1000`
  - `GITEA__database__DB_TYPE=mysql`
  - `GITEA__database__HOST=mysql:3306`
  - `GITEA__database__NAME=gitea`
  - `GITEA__database__USER=gitea`
  - `GITEA__database__PASSWD=gitea`

If you started the container you can finish the setup of gitea in a browser at [http://localhost:3000](http://localhost:3000). All settings can be left blank.

Note: once you hit the install button, it takes around a minute for gitea to install itself.

## Bonus

Try to create an inital admin user, create a repo and try to clone it.

What happens when the container is stopped? What happens when you remove it?

## Solution

See [here](./solution.md)
