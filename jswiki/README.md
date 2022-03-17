# JSWiki

The following command starts a jswiki container using a local sqlite file as persistency.

```bash
docker run -d -p 8080:3000 --name wiki --restart unless-stopped -e "DB_TYPE=sqlite" -e "DB_FILEPATH=/wiki/wiki.sqlite" -v "/home/technat/code/docker-ict-campus-demo/jswiki/wiki.sqlite:/data/wiki.sqlite" ghcr.io/requarks/wiki:2
```
