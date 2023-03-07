# Filebrowser

[Filebrowser](https://filebroser.org) is a simple web-based filebrowser (file explorer) that you can use to share files with others or access volumes of other containers easily.

To run the container you need the following:

- Port `80` exposed to the host as port `8080`
- Run in detached mode
- Name the container `filebrowser`
- Mount a locally created directory `./data` inside the container at `/src`
- Mount the local file (which you must create first) `./filebrowser.db` inside the container at `/database.db`
- Use image `filebrowser/filebrowser`

Consult their [docs](https://filebrowser.org/installation#docker) for more infos.

If everything was successful you can open the web-UI at [http://localhost:8080](http://localhost:8080) and login with `admin` and PW `admin`.

## Bonus

Instead of the locally created `./data` directory use a [docker volume](https://docs.docker.com/storage/volumes/).

## Solution

See [here](./solution.md)
