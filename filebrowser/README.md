# Filebrowser

The following commands create a docker volume to serve files from, mount a DB file into the container and start it, exposing the UI on `http://localhost:8080`.

```bash
docker volume create filebrowser
touch filebrowser.db
docker run -d --name filebrowser -v /home/technat/code/docker-ict-campus-demo/filebrowser/filebrowser.db:/database.db -p 8080:80 --mount source=filebrowser,target=/src filebrowser/filebrowser
```

Default Admin: `admin` and `admin`
