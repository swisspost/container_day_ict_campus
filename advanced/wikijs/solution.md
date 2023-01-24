# JSWiki

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

The only file that needs to be backed up is the wiki.sqlite file you mounted into the container. All data within the container is ephemeral and can be treated like that.

If you run the same command on another docker host, with the wiki.sqlite file copied there first, the wiki starts just as you would expect.
