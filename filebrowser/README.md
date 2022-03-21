# Filebrowser

Change to the filebrowser directory:

```bash
cd filebrowser
```

Create a new docker-volume:

```bash
docker volume create filebrowser
```

Create the Database file in this directory:

```bash
touch filebrowser.db
```

Start the container:

```bash
docker run -d --name filebrowser -v <PATH to Repo>/filebrowser/filebrowser.db:/database.db -p 8080:80 --mount source=filebrowser,target=/src filebrowser/filebrowser
```

You can find the UI at [http://localhost:8080](http://localhost:8080). PW: `admin`, User: `admin`

Stop and delete the running container:

```bash
docker stop <container_id>  
docker rm <container_id> 
```
