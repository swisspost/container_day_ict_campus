# Webserver

This exercise is about building a custom image using a `Dockerfile`.

Write a new dockerfile that:
- uses `alpine:latest` as base
- installs `httpd` as webserver
- creates a workding directory in `/var/www/localhost/htdocs`
- copies the index.html into this directory
- ensures the correct permissions (755) on the index.html
- sets the entrypoint directive to start `httpd`
- sets the CMD directive to run in foreground
- exposes port 80

Then build the dockerfile to an image named `webserver:latest` and run it so that [http://localhost:8080](http://localhost:8080) displays your simple HTML page.

## Solution

See the [Dockerfile](./Dockerfile) for a working example and [here](./solution.md) for the commands to build and run it.