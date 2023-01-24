# Filebrowser

First create the empty DB file:

```bash
touch ./filebrowser.db
```

And the data directory:

```bash
mkdir -p ./data
```

To run the filebrowser you need this command:

```bash
docker run -d \
  --name filebrowser \
  -p 8080:80 \
  -v <PATH to current directory>/filebrowser/filebrowser.db:/database.db \
  -v <Path to current directory>/data:/src
  --mount source=filebrowser,target=/src \
  filebrowser/filebrowser
```

The options in detail:

- `-d`: Run in detached mode (you could close your shell and the filebrowser should continue to run)
- `--name filebrowser`: With a name it's easier to stop/start the container later
- `-p 8080:80`: exposes port 8080 on the host and redirect to the container's port 80
- `-v <Path to the current directory>/wiki.sqlite:/wiki/wiki.sqlite`: Mounts the file database.db from the current directory into the container at `/database.db`
- `filebrowser/filebrowser:latest`: The image that is started

Stop and delete the running container:

```bash
docker stop filebrowser
docker rm filebrowser
```

## Bonus

To create a new docker-volume run the following:

```bash
docker volume create filebrowser
```

And use the following parameter in the docker-run command:

```bash
--mount source=filebrowser,target=/src
```

e.g:

```bash
docker run -d \
  --name filebrowser \
  -p 8080:80 \
  -v <PATH to current directory>/filebrowser/filebrowser.db:/database.db \
  --mount source=filebrowser,target=/src \
  filebrowser/filebrowser
```
