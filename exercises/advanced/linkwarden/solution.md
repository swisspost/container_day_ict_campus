# Linkwarden - solution

## Network

Create the bridge network `linkwarden`:

```bash
docker network create -d bridge linkwarden
```

## Database

To run the postgresql you need this command:

```bash
docker run -d \
  --name linkwarden-db \
  --network linkwarden \
  -v <Path to current directory>/pgdata:/var/lib/postgresql/data \
  -e POSTGRES_PASSWORD="XTA25bhckq9G4kGHJf9n5CNhrCcHdEvv" \
  postgres:16-alpine
```

The options in detail:

- `--network linkwarden`: uses dedicated bridge network for linkwarden
- `-v <PATH to current directory>/pgdata:/var/lib/postgresql/data`: mounts the local directory to the container at the mounting point /var/lib/postgresql/data
- `-e POSTGRES_PASSWORD="XTA25bhckq9G4kGHJf9n5CNhrCcHdEvv"`: sets environment variable for the database password
- `postgres:16-alpine`: uses postgresql 16 docker image based on alpine

## Linkwarden

As soon as the database is running, you need this command in order to run linkwarden:

```bash
docker run -d \
  --name linkwarden \
  --network linkwarden \
  -p 3000:3000 \
  -v <Path to current directory>/data:/data/data \
  -e NEXTAUTH_URL="http://localhost:3000/api/v1/auth" \
  -e NEXTAUTH_SECRET="pbRN6im6Z9kbLq5tTHxfD78KtVM9SroH" \
  -e DATABASE_URL=postgresql://postgres:XTA25bhckq9G4kGHJf9n5CNhrCcHdEvv@linkwarden-db:5432/postgres \
  ghcr.io/linkwarden/linkwarden
```

The options in detail:

- `--network linkwarden`: uses dedicated bridge network for linkwarden
- `-p 3000:3000`: exposes port 3000 on the host and redirect to the container's port 3000
- `-v <PATH to current directory>/data:/data/data`: mounts the local directory to the container at the mounting point /data/data
- `-e NEXTAUTH_URL="http://localhost:3000/api/v1/auth"`: sets environment variable for the authentication url
- `-e NEXTAUTH_SECRET="pbRN6im6Z9kbLq5tTHxfD78KtVM9SroH"`: sets environment variable for the authentication secret
- `-e DATABASE_URL=postgresql://postgres:XTA25bhckq9G4kGHJf9n5CNhrCcHdEvv@linkwarden-db:5432/postgres`: sets environment variable for the database connection string with the user postgres, password from before, postgresql server reachable at linkwarden-db:5432 and the database postgresql
- `ghcr.io/linkwarden/linkwarden`: uses latest docker image of linkwarden


## Bonus

Headover to the [documentation](https://docs.linkwarden.app/self-hosting/environment-variables) and search for the enviornment variables. There you can see that the environment variable needs to be set to `NEXT_PUBLIC_DISABLE_REGISTRATION=true`. 


## Clean up

Stop and delete the running container:

```bash
docker stop linkwarden
docker stop linkwarden-db
docker rm linkwarden
docker rm linkwarden-db
```
