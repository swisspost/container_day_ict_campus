# Linkwarden

[Linkwarden](https://linkwarden.app/) is a collaborative bookmark manager to collect, organize, and preserve webpages and articles. 

## Network

Create a new bridge network `linkwarden` for this exercise.

## Database

Linkwarden depends on a postgresql database. Therefore you need to set it up first:

- Mount a locally created directory `./pgdata` inside the container at `/var/lib/postgresql/data`
- Set environment variable POSTGRES_PASSWORD
- Name the container `linkwarden-db`
- Use image `postgres:16-alpine`

## Linkwarden

To run the application you need the following:

- Expose port `3000` to the host as port `3000`
- Mount a locally created directory `./data` inside the container at `/data/data`
- Set environment variable NEXTAUTH_URL to `http://localhost:3000/api/v1/auth`
- Set environment variable NEXTAUTH_SECRET
- Set environment variable DATABASE_URL `postgresql://postgres:<POSTGRES_PASSWORD>@linkwarden-db:5432/postgres` and replace the `<POSTGRES_PASSWORD>` with the password set for the postgresql database 
- Name the container `linkwarden`
- Use image `ghcr.io/linkwarden/linkwarden`


If everything was successful you can open the web-UI at [http://localhost:3000](http://localhost:3000) and create a new user.


## Bonus

After you logged in, disable the registration. Check the [documentation](https://docs.linkwarden.app).

## Solution

See [here](./solution.md)
