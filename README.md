# Coconut

Coconut is a starter template for building web servers in Go.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/TropiStack/coconut.git
   cd coconut
   ```

2. **Install dependencies:**

   Ensure you have Go installed, then run:

   ```bash
   go mod tidy
   ```

3. **Build the project:**

   ```bash
   go build -o coconut main.go
   ```

## Usage

### Running Locally

To start the server locally, run:

```bash
./coconut
```

The server will start on `http://localhost:8080` by default.

### Using Docker

#### Development

To run the server with live reloading using Docker, use the following command:

```bash
docker-compose -p coconut-dev -f docker/docker-compose.yml -f docker/docker-compose.dev.yml up
```

#### Production

To run the server in a production environment, use:

```bash
docker-compose -p coconut-prod -f docker/docker-compose.yml -f docker/docker-compose.prod.yml up
```

## Contributing

Currently, we are not accepting outside contributions. We appreciate your interest and encourage you to check back in the future for any updates regarding contribution opportunities.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
