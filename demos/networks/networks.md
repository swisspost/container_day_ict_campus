# Commands
## Networks
```bash
docker network ls
```

## Port-Forwarding
```bash
docker run --rm -d -p 8080:80 --name nginx nginx
```
```bash
docker inspect nginx | jq '.[].NetworkSettings.Ports'
```

## Custom Network
```bash
docker network create mynet --subnet 10.11.12.0/24
```
```bash
docker run --rm -d --network mynet -p 8080:80 --name nginx nginx
```
```bash
docker inspect nginx | jq '.[].NetworkSettings.Networks'
```

## DNS-File
```bash
docker run --rm -d -p 8080:80 --name nginx nginx
```
```bash
cat /etc/resolv.conf
```
```bash
docker exec -ti nginx cat /etc/resolv.conf
```

## DNS-Server
```bash
docker network create asdf
docker network create gugus
```
```bash
docker run --rm -d --network=asdf --name asdf-server nginx
```
```bash
docker run --rm --network=asdf alpine/curl asdf-server
```
```bash
docker run --rm --network=gugus alpine/curl asdf-server
```
