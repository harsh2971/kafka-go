# Kafka Go - Order Management System

A real-time order management system built with Go and Apache Kafka, demonstrating producer-consumer patterns with Docker containerization.

## 🚀 Features

- **Real-time Message Processing**: Producer-Consumer architecture using Apache Kafka
- **Web Interface**: Simple HTML form to submit orders
- **Docker Compose Setup**: Complete containerized environment with Kafka, Zookeeper, and the Go application
- **Automatic Topic Creation**: Kafka topics are created programmatically on startup
- **Live Logging**: Real-time order processing logs visible in the console

## 📋 Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop) installed and running
- [Go 1.23+](https://go.dev/dl/) (for local development)
- Git

## 🏗️ Project Structure

```
kafka-go/
├── api/
│   └── main.go              # Main application entry point
├── orders/
│   ├── consumer.go          # Kafka consumer implementation
│   ├── producer.go          # Kafka producer implementation
│   ├── createTopic.go       # Topic creation logic
│   ├── handlers.go          # HTTP request handlers
│   └── model.go             # Order data model
├── web/
│   └── index.html           # Order submission form
├── docker-compose.yml       # Docker services configuration
├── Dockerfile               # Multi-stage Docker build
├── go.mod                   # Go module dependencies
└── README.md
```

## 🛠️ Installation & Setup

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/harsh2971/kafka-go.git
   cd kafka-go
   ```

2. **Start all services**
   ```bash
   docker compose up --build
   ```

3. **Access the application**
   - Open your browser: http://localhost:8080
   - Fill out the order form and submit

### Local Development

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Start Kafka and Zookeeper**
   ```bash
   docker compose up zookeeper kafka -d
   ```

3. **Run the application locally**
   ```bash
   cd api
   go run main.go
   ```

## 📊 Architecture

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│             │         │              │         │             │
│  Web Client │────────▶│  HTTP Server │────────▶│   Producer  │
│             │         │              │         │             │
└─────────────┘         └──────────────┘         └──────┬──────┘
                                                         │
                                                         ▼
                                                  ┌─────────────┐
                                                  │    Kafka    │
                                                  │   Broker    │
                                                  └──────┬──────┘
                                                         │
                                                         ▼
                                                  ┌─────────────┐
                                                  │  Consumer   │
                                                  │  (Logs)     │
                                                  └─────────────┘
```

## 🔧 Configuration

### Environment Variables

The application uses the following environment variable:

- `KAFKA_BROKER`: Kafka broker address (default: `localhost:29092`)

### Docker Services

- **Zookeeper**: Port 2181
- **Kafka**: 
  - External: Port 29092 (for host machine)
  - Internal: Port 9092 (for Docker network)
- **API Server**: Port 8080

## 📝 Usage

### Submit an Order

1. Navigate to http://localhost:8080
2. Fill in the order form:
   - Product name (e.g., "t-shirt")
   - Quantity (e.g., 5)
3. Click "Submit Order"

### View Logs

**Follow all service logs:**
```bash
docker compose logs -f
```

**Follow only API logs:**
```bash
docker compose logs -f api
```

**Example log output:**
```
api  | Topic created
api  | Server started on port 8080
api  | Consumer listening for orders
api  | Producer sent the order to Kafka
api  | Received order: t-shirt 55
```

## 🔍 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/`      | Serve the order form HTML page |
| POST   | `/order` | Submit a new order (JSON body) |

### Example Request

```bash
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{
    "product": "laptop",
    "quantity": 2
  }'
```

## 🛑 Stopping the Application

```bash
# Stop and remove containers
docker compose down

# Stop and remove containers with volumes
docker compose down -v
```

## 🧪 Testing

1. **Check service status:**
   ```bash
   docker compose ps
   ```

2. **Verify Kafka topics:**
   ```bash
   docker exec -it kafka kafka-topics --list --bootstrap-server localhost:9092
   ```

3. **Monitor messages in real-time:**
   ```bash
   docker compose logs -f api
   ```

## 🐛 Troubleshooting

### Port Already in Use

If you get a "port already in use" error:
```bash
# Find and kill the process using port 8080
lsof -ti:8080 | xargs kill -9

# Or stop Docker containers
docker compose down
```

### Kafka Connection Issues

If the API can't connect to Kafka:
1. Ensure Kafka is healthy: `docker compose ps`
2. Wait for Kafka health check to pass (30 seconds)
3. Check logs: `docker compose logs kafka`

### Docker Build Fails

```bash
# Clean build cache and rebuild
docker compose down
docker system prune -a
docker compose up --build
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🔗 Resources

- [Apache Kafka Documentation](https://kafka.apache.org/documentation/)
- [segmentio/kafka-go](https://github.com/segmentio/kafka-go) - Kafka client library for Go
- [Docker Compose Documentation](https://docs.docker.com/compose/)

## 👤 Author

**harsh2971**
- GitHub: [@harsh2971](https://github.com/harsh2971)

---

⭐ Star this repository if you find it helpful!

