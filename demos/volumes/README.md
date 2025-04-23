# Volumes to persist data


```console
docker run -ti --rm -p 8080:8080 --name web webserver # uses html from container
docker run -ti --rm -p 8080:8080 --name web -v ./html:/usr/share/nginx/html webserver # uses html from host, edits are live-synced without a container restart
```
