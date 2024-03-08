# Pi-hole

[Pi-hole](https://pi-hole.net/) is a network-level advertisement and internet tracker blocking application, that enhances privacy and security by intercepting and blocking DNS requests for known advertising and tracking domains.

To run the container you need the following:

- Expose port `80` to the host as port `8080`
- Expose port `53/udp` to the host as port `1053/udp`
- Expose port `53/tcp` to the host as port `1053/tcp`
- Set timezone to `Europe/Zurich` with the environment variable `TZ`
- Set admin password with the environment variable `WEBPASSWORD`
- Name the container `pihole`
- Mount a locally created directory `./pihole` inside the container at `/etc/pihole`
- Mount a locally created directory `./dnsmasq.d` inside the container at `/etc/dnsmasq.d`
- Add capabilities `NET_ADMIN`
- Use image `pihole/pihole`

If everything was successful you can open the web-UI at [http://localhost:8080](http://localhost:8080) and login with the PW you set to the environment variable.

## Bonus

Send a dns request to the pihole which:
- gets resolved
- gets blocked

## Solution

See [here](./solution.md)
