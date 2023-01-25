# JSWiki

[WikiJS](https://js.wiki) is a powerful open source wiki software written using NodeJS. It can easily be run as a standalone container.

Run a new wikijs container accoring to their [docs](https://docs.requarks.io/install/docker) with the following options:

- Name the container `wiki`
- set ENV var `DB_TYPE=sqlite`
- set ENV var `DB_FILEPATH=/wiki/wiki.sqlite`
- mount the file ./wiki.sqlite (use absolut path) into /wiki/wiki.sqlite in the container (file will be created automatically if it doesn't exist)
- expose port 3000 as 3000
- run in detached mode
- Use image `ghcr.io/requarks/wiki:2`

If everything worked successfully you can open [http://localhost:3000] and finish the setup of the wiki.

## Bonus

What is needed to run your installed wiki on another host? Which files would you need to backup?

## Solution

See [here](./solution.md)
