# JSWiki

[WikiJS](https://js.wiki) is a powerful open source wiki software written using NodeJS. It can easily be run as a standalone container.

Run a new wikijs container accoring to their [docs](https://docs.requarks.io/install/docker) with the following options:

- Name the container `wiki`
- set ENV var `DB_TYPE=sqlite`
- set ENV var `DB_FILEPATH=/wiki/wiki.sqlite`
- mount the file `./wiki.sqlite` (use absolut path) as `/wiki/wiki.sqlite` in the container (file will be created automatically if it doesn't exist)
- expose port `3000` as `3000`
- run in detached mode
- Use image `ghcr.io/requarks/wiki:2`

If everything worked successfully you can open [http://localhost:3000] and finish the setup of the wiki.

## Bonus

What is needed to run your installed wiki on another host? Which files would you need to backup?

## Bonus 2

Try to setup another high-available wikijs instance with a PostgresSQL DB instead of sqlite

Some hints:

- You need to put the containers in a separate network (`docker network create wikijs_postgres`, `--net wikijs_postgres`)
- You need to run a postgres container (the internet has thousand examples)
- Spin up three wikijs instances for HA
- You need to configure wikijs to use the postgres database (see [here](https://docs.requarks.io/install/docker))
- You need to set the env var `HA_ACTIVE=true` on the wikijs containers
- Expose the containers on 3001,3002 and 3002

## Super bonus

If you have already done the [webserver-ha](../webserver-ha/) example, you may know what a reverse proxy is. If not go and do this example first. If you have already done it, do the following:

- Copy the [Caddyfile](../webserver-ha/Caddyfile) to this directory
- Edit the file so that it forwards to the three wiki containers and the correct ports
- Run a new reverse_proxy container and expose it on port 80 and put it in the correct network ;)
- Check if [http://localhost:80](http://localhost:80) shows the wiki

## Solution

See [here](./solution.md)
