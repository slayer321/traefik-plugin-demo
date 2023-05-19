#### Traefik-Plugin 

How to Run it on local 

1. First run the whoami container

```
docker run -d --network host traefik/whoami -port 5000
```

2. Run the traefik server
```
git clone https://github.com/slayer321/traefik-plugin-demo.git
cd traefik-plugin-demo
traefik --configfile traefik-static.yaml
```