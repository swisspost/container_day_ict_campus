# Pi-hole - solution

Create the directories:

```bash
mkdir -p ./pihole
mkdir -p ./dnsmasq.d
```

To run the pihole you need this command:

```bash
docker run -d \
  --name pihole \
  -p 8080:80 \
  -p 1053:53/udp \
  -p 1053:53/tcp \
  -e TZ="Europe/Zurich" \
  -e WEBPASSWORD="pihole" \
  -v <PATH to current directory>/pihole:/etc/pihole \
  -v <Path to current directory>/dnsmasq.d:/etc/dnsmasq.d \
  --cap-add=NET_ADMIN \
  pihole/pihole:latest
```

The options in detail:

- `-p 8080:80`: exposes port 8080 on the host and redirect to the container's port 80
- `-p 1053:53/udp`: exposes port 1053/udp on the host and redirect to the container's port 53/udp
- `-p 1053:53/tcp`: exposes port 1053/tcp on the host and redirect to the container's port 53/tcp
- `-e TZ="Europe/Zurich`: sets environment variable for the timezone
- `-e WEBPASSWORD="pihole"`: sets environment variable for the admin password used to login to the webui
- `-v <PATH to current directory>/pihole:/etc/pihole`: mounts the local directory to the container at the mounting point /etc/pihole
- `-v <Path to current directory>/dnsmasq.d:/etc/dnsmasq.d`: mounts the local directory to the container at the mounting point /etc/dnsmasq.d
- `--cap-add=NET_ADMIN`: adds the capability to modify network settings
- `pihole/pihole:latest`: uses latest docker image of pihle

Stop and delete the running container:

```bash
docker stop pihole
docker rm pihole
```

## Bonus

On linux we can use a cli tool like dig to send a dns request to the pihole

```bash
dig -p 1053 @localhost google.ch
```

Now lets check if it actually blocks ads

```bash
dig -p 1053 @localhost googleads.g.doubleclick.net 
```