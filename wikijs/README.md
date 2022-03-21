# JSWiki

Start a jswiki container using the following command: 

```bash
docker run -d -p 3000:3000 --name wiki -e "DB_TYPE=sqlite" -e "DB_FILEPATH=/wiki/wiki.sqlite" -v "<PATH to Repo>/jswiki/wiki.sqlite:/wiki/wiki.sqlite" ghcr.io/requarks/wiki:2
```

To finish the setup head over to [http://localhost:3000](http://localhost:3000) to finish the setup.

Stop and delete the running container:

```bash
docker stop <container_id>  
docker rm <container_id> 
```
