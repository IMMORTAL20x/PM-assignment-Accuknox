# DateTime Web Application with Kubernetes Deployment

## Project Overview
A web application built with Go that displays current date and time, containerized with Docker and deployed on Kubernetes using Minikube.

## Tech Stack
- Backend: Go
- Containerization: Docker
- Orchestration: Kubernetes
- Local Development: Minikube

## Project Structure
```
.
├── src/
│   └── main.go                 # Go application source code
├── kubernetes/
│   ├── deployment.yaml         # Kubernetes deployment configuration
│   └── service.yaml           # Kubernetes service configuration
├── docker/
│   └── Dockerfile             # Docker build configuration
├── docs/
│   ├── screenshots/           # Implementation screenshots
│   └── PRD.md                # Product Requirements Document
└── README.md                 # This file
```

## Implementation Steps

1. **Go Application Development**
   - Created web server displaying date and time
   - Implemented HTTP endpoints
   - Tested locally

2. **Docker Containerization**
   - Created Dockerfile
   - Built and tested container
   - Pushed to Docker Hub

3. **Kubernetes Deployment**
   - Created deployment with 2 replicas
   - Set up LoadBalancer service
   - Deployed using Minikube
   - Exposed service using tunnel

## Deployment Instructions

1. **Start Minikube**
```bash
minikube start
```

2. **Apply Kubernetes Configurations**
```bash
kubectl apply -f kubernetes/deployment.yaml
kubectl apply -f kubernetes/service.yaml
```

3. **Access Application**
```bash
minikube tunnel
# Access via http://localhost:80
```

## Screenshots
- Application running locally
- Docker build process
- Kubernetes deployment
- Final application access

## Documentation
Detailed documentation available in:
- `docs/PRD.md` - Product Requirements Document
- `docs/screenshots/` - Implementation screenshots

## Current Status
✅ Go Application Development
✅ Docker Container Creation
✅ Kubernetes Deployment
✅ Service Exposure
✅ Documentation

## Author
IMMORTAL20x
