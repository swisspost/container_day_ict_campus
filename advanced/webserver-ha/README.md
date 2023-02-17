# Webserver HA

In this exercise you are going to deploy your webserver in high-available mode using a reverse_proxy in front of it.

## Step 1 - make color of website customizable

To determine which of your multiple webservers responds to your request, we want to show a different page for each of our webservers. To do this, we create three different index.html files that can be mounted inside the containers.

For each file, choose another background color and save them as `index1.html`, `index2.html` and `index3.html`

## Step 2 - Create a docker network

We haven't talked about [docker networks](https://docs.docker.com/network/) so far, but think of them as separate virtual networks, just as you got multiple subnets in your virtualized environment. They are used to separate multiple containerized apps from each other and have the benefit that containers can directly resolve each other by their hostname (that's why it's important to give them a name using `--name`).

For this app we are going to create our own network using `docker network create webserver-ha`. 

You can find a list of all networks using `docker network ls`.

To start a container in a separate network, run it using the `--net webserver-ha` option.

## Step 3 - Start multiple webservers

Now that we have three different websites, start three different webservers, for each mount the index{1,2,3}.html as `/var/www/localhost/htdocs/index.html` in the container.

**Don't expose the port as it would cause conflicts! Instead, put the container in the created docker network.**

If you have them started, you should see the following in the `docker ps` output:

```
CONTAINER ID   IMAGE        COMMAND                  CREATED          STATUS                  PORTS   NAMES
5fb519969a95   webserver    "/usr/sbin/httpd -D …"   2 seconds ago    Up Less than a second   80/tcp  webserver3
69ce51f3cc2f   webserver    "/usr/sbin/httpd -D …"   11 seconds ago   Up 9 seconds            80/tcp  webserver2
bfee8d63284b   webserver    "/usr/sbin/httpd -D …"   20 seconds ago   Up 18 seconds           80/tcp  webserver1
```

## Step 4 - Start a loadbalancer

Now you might notice that we haven't exposed any ports on these containers to the host. This is because we can't map all of them to port 8080 as port 8080 is unique. Of course we could map them to different ports, but then we would have three different URLs. To distribute the traffic now and expose the service securly we are using a reserve_proxy or depending on the layer also called loadbalancer.

If you want to know more about reserve proxies, checkout [this explanation](https://www.cloudflare.com/learning/cdn/glossary/reverse-proxy/).

We are using [caddy](https://caddyserver.com/) here as reserve proxy since the author of this exercise prefers caddy over nginx or apache which could also function as reserve proxies. The config for the reserve proxy is already written, see the [Caddyfile](./Caddyfile) if you are interested.

Now to start the caddy as reverse proxy, run a new container with:
- detached mode
- network: webserver-ha
- name `caddy`
- expose port 80 as 8080 on the host
- mount the `./Caddyfile` into `/etc/caddy/Caddyfile` in the container

If the caddy is started, you should be able to browse [http://localhost:8080](http://localhost:8080) and see a different page on every refresh.
