# Gitea

Change to the gitea directory:

```bash
cd gitea
```

Adjust the `docker-compose.yml` file and replace `USER_UID` and `USER_GID` with the appropriate output of `id`.

Then run:

```bash
docker-compose up -d
```

Finish the installation of Gitea in the brower at [http://localhost:3000](http://localhost:3000).

Stop and delete the containers using:
```bash
docker-compose down
```
