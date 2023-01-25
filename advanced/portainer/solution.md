# Portainer - solution

## Step 1 - Create docker volume

```bash
docker volume create portainer_data
```

## Step 2 - Run Container

```bash
docker run -d -p 8000:8000 -p 9443:9443 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce:latest
```

Access the UI at [https://localhost:9443](https://localhost:9443), finish the Setup and explore the UI ;).
