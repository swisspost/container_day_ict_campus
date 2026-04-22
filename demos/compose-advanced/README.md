# Docker Compose Advanced Example

This deploys a Mattermost Chat Server as Docker Compose Stack

## Deploy

Create a `.env` file with the following values:
- `POSTGRES_USER`: user for the db
- `POSTGRES_PASSWORD`: password for the DB user
- `POSTGRES_DB`: name of the DB to use
- `DOMAIN`: e.g localhost:443 or similar if you exposed Caddy on this port or a domain if you have some DNS records

And then run:

```console
docker compose up -d
```

Open the DOMAIN endpoint in your browser to finish the setup.
