# Container Day ICT-Campus

Introduction to containers for trainees of SwissPost

## About

The container day is a one-day workshop that teaches trainees in their frist year of apprenticeship the basics about containers. The goal in the end of this workshop is, that everyone is comfortable building their own images and running containers using Docker on a single host.

There's also a short introduction to Kubernetes to show that there's more than just plain Docker.

## Help

- [Docker Cheat Sheet](https://dockerlabs.collabnix.com/docker/cheatsheet/)

## Exercises

Those are exercises that guide you through running pre-built and self-built containers using docker. It's recommended to go through the exercises in the order they are listed below.

### Prepare

To run the exercises you must have:

- cloned the repo
- docker [installed](https://docs.docker.com/engine/install/ubuntu/)
- an editor installed

To start with, [installing](https://docs.portainer.io/start/install-ce/server/docker/linux) Portainer can help:
```bash
docker volume create portainer_data 
docker run -d -p 8000:8000 -p 9000:9000 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce:latest
```

Now you can visit the admin dashboard in the browser:`http://localhost:9000`

**Note: for every exercise it's important that you change inside it's directory using `cd`, otherwise the solutions won't work.**

### Beginner

- [Hello World](./beginner/helloworld): Very basic dockerfile that runs a bash script at startup (same as in live demo)
- [Webserver](./beginner/webserver): Very simple webserver that presents a static website

### Advanced

- [Highly available webserver](./advanced/webserver-ha): High-available version of the static website using multiple webservers and a reverse-proxy
- [filebrowser](./advanced/filebrowser/): Basic web-based filebrowser to view/edit/share files on the web
- [Gitea](./advanced/gitea): Lightweight but powerful git server to self-host as container (uses a mysql database)
- [wikijs](./advanced/wikijs/): Wonderful wiki solution for a self-hosted wiki
- [pihole](./advanced/pihole/): Advertisement and internet tracker blocking application
- [linkwarden](./advanced/linkwarden/): Collaborative bookmark manager

### Expert

If you want to go deeper into container and kubernetes checkout the following links:

- [Docker-Compose guide](https://gabrieltanner.org/blog/docker-compose/): A guide showing how to use docker-compose to define container configuration using YAML
- [Add more docker hosts to portainer](https://markontech.com/devops/add-a-remote-docker-host-in-portainer/): Scaling using multiple docker hosts managed by a single user-interface
- [How to write a best-practise Dockerfile](https://sysdig.com/blog/dockerfile-best-practices/): To study, those are best-practices everyone should follow but 90% of containers don't (which makes them very insecure)
- [Demystifying Containers](https://github.com/saschagrunert/demystifying-containers/tree/master): A good interactive guide explaining how containers work under the hood
- [Quickstart on how to setup a k3s cluster](https://docs.k3s.io/quick-start): Docs how to setup a lightweight Kubernetes cluster (wait for the K8s intro first)
- [Setup Kubernetes the kubeadm way](https://technat.ch/posts/k8s_kubeadm/): Searching something to do in your project week? This could be your project! Setting up kubernetes using kubeadm is time-intensive but helps you really learn how kubernetes works. Only recommended for people that want to go deep into the topic
