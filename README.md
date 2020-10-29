# go_webapp

# Prerequisites

docker
docker-compose
golang

# Deployment
## Build Docker image for production environment

```bash
git clone https://github.com/hungdoo/go_webapp 
cd go_webapp/ && docker build -t go-webapp-prod -f Dockerfile.prod .
```

## Run container as daemon

docker run -it -p 80:8080 -d go-webapp-prod

# Test
cd src/
go test -v . ./api/
go test -cover . ./api/