# Solution for Docker Registry.

There's an official image from docker to host your own registry: https://hub.docker.com/_/registry

```console
docker run \
  -d \
  -p 5000:5000 \
  --restart always \
  --name registry \
  registry:3
```

You can push and pull images to this registry using:
```console
docker tag myimage localhost:5000/myimage
docker push localhost:5000/myimage
docker pull localhost:5000/myimage
```

## Pull-through cache

```console
## using docker-cli
docker run \
  -d \
  -p 5000:5000 \
  --restart always \
  --name registry \
  -e REGISTRY_PROXY_REMOTEURL="https://registry-1.docker.io" \
  registry:3

## using docker compose
docker compose up -d
```

Configure the local docker daemon like described [here](https://distribution.github.io/distribution/recipes/mirror/#configure-the-docker-daemon):

```console
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "registry-mirrors": ["https://localhost:5000"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
docker info | grep "mirror" -i -A 1
```

### Test 

Pull an image on your host and check the logs of your registry (`docker logs registry`) to see that it will pull and serve you the image you requested.
