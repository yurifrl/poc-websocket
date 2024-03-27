
# WebSocket Server POC

This guide outlines the steps for setting up and running the WebSocket server and client locally for development purposes.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/) cluster (e.g., [Minikube](https://minikube.sigs.k8s.io/docs/start/), [Kind](https://kind.sigs.k8s.io/))
- [Skaffold](https://skaffold.dev/)
- [Chaos Mesh](https://chaos-mesh.org/) (for chaos testing)
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)

## Setup

1. **Start your local Kubernetes cluster:**

   For Minikube:

   ```bash
   minikube start
   ```

   For Kind:

   ```bash
   kind create cluster
   ```

2. **Install Chaos Mesh:**

   Follow the [Chaos Mesh installation guide](https://chaos-mesh.org/docs/user_guides/installation) to install Chaos Mesh in your local Kubernetes cluster.

4. **Start the WebSocket server:**

   Use Skaffold to deploy the WebSocket server to your local Kubernetes cluster:

   ```bash
   skaffold dev --tail --trigger=manual
   ```

   This command starts the WebSocket server in development mode with manual trigger for rebuilds and tails the logs.

5. **Run the WebSocket client:**

   In a separate terminal, run the WebSocket client:

   ```bash
   go run cmd/client/*.go exec
   ```

   This command executes the client, which connects to the WebSocket server and sends/receives messages.


## To Generate Protocol Buffers code

   Navigate to the project root directory and run:

   ```bash
   protoc --go_out=. --go_opt=paths=source_relative proto/message.proto
   ```

   Alternatively, using Docker:

   ```bash
   docker compose run --rm app protoc --go_out=. --go_opt=paths=source_relative proto/message.proto
   ```

## Additional Notes

- You can use Chaos Mesh to perform chaos experiments on your WebSocket server to test its resilience and fault tolerance.
- Make sure to configure your Kubernetes cluster and Skaffold settings according to your development environment.