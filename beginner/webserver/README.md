# Webserver

This exercise is about building a custom image that serves a static website.

Write a new Dockerfile that:

- uses `alpine:latest` as base
- installs `nginx` as webserver
- copies the nginx.conf into `/etc/nginx/`
- copies the index.html into `/usr/share/nginx/html/`
- sets the ENTRYPOINT directive to start `nginx`
- sets the CMD directive to run in foreground
- exposes port 8080

Then build the Dockerfile to an image named `webserver:latest` and run it so that [http://localhost:8080](http://localhost:8080) displays your simple HTML page.

## Bonus

What if instead of our static `index.html` you would add your self-made webpage from the web module into the container?

## Solution

See the [Dockerfile](./Dockerfile) for a working example and [here](./solution.md) for the commands to build and run it.
