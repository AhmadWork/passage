# Passage - Header-Based Reverse Proxy Load Balancer

**Passage** is a lightweight, header-based reverse proxy written in Go. It allows traffic routing to different backends based on a custom HTTP header (`App-Name`). This makes it a simple and effective solution for managing multiple backend services running under the same URL, such as microservices deployed behind a single entry point.

---

## ✨ Features

- **Header-Based Routing:** Uses `App-Name` header to decide which backend to forward the request to.
- **Health Checks:** Periodic health checks ensure traffic is only sent to healthy backends.
- **Automatic Retry:** Automatically retries failed requests up to 3 times before marking a backend as down.
- **Fallback Support:** If the specified `App-Name` backend is unavailable, traffic can fallback to a default backend.
- **Docker Ready:** Comes with a Dockerfile and Docker Compose setup for easy local testing.

---

## 🏗️ Architecture

```
            Client Request
                    |
              +----------------+
              |   Passage LB    |
              +----------------+
                    |
    +-------------------------------+
    |        Backend Selection      |
    |    (based on App-Name header) |
    +-------------------------------+
         |            |          |
      backend1    backend2    backend3
```

---

## 🚀 Quick Start

### Clone the repository

```bash
git clone https://github.com/your-username/passage.git
cd passage
```

### Build & Run

```bash
go build -o passage
./passage --backends "fms-http://backend:8001,lms-http://app:80,http://backend:8000" --port 3030
```

### Example Request

```http
GET /path HTTP/1.1
Host: localhost:3030
App-Name: fms
```

---

## 🐳 Running with Docker Compose

Passage comes with a ready-to-use `docker-compose.yml` file.

```bash
docker-compose up --build
```

This will:

- Build and run the Passage container.
- Start three sample backend services using `strm/helloworld-http` images.
- Expose Passage at `http://localhost:3030`.

You can send requests with different `App-Name` headers to see traffic routed to different backends.

---

## ⚙️ Configuration

### CLI Flags

| Flag          | Description                                                | Example |
|---------------|------------------------------------------------------------|---|
| `--backends`  | Comma-separated list of backends. Each can have a key-prefix like `fms-http://backend:8001` | `fms-http://backend:8001,lms-http://app:80,http://backend:8000` |
| `--port`      | Port for Passage to listen on                              | `--port 3030` |

---

## 📥 App-Name Header

The `App-Name` header controls which backend handles a request.

| Header Value | Target Backend |
|---|---|
| `fms`  | `http://backend:8001` |
| `lms`  | `http://app:80` |
| *(any other value)* | `http://backend:8000` (default) |

---

## 📊 Health Checks

- All backends are checked every **2 minutes**.
- Backends marked as **down** will not receive traffic.
- Health check is a simple **TCP connection test**.

---

## 🛠️ Internal Structure

| Component     | Purpose |
|---|---|
| `ServerPool`  | Manages available backends and their health status. |
| `Backend`     | Represents a backend service with its URL, name, and reverse proxy. |
| `crw`         | A custom response writer used to log response data. |

---

## 📂 Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── main.go (the core logic)
├── internal/
│   └── crw/
│       └── logging_response_writer.go
├── go.mod
└── go.sum
```

---

## 🪪 License

This project is licensed under the **MIT License** — meaning you're free to copy, modify, distribute, and use it in both personal and commercial projects.

---

## 👥 Contributing

Contributions are welcome! Feel free to open an issue or submit a PR.

---

## 📧 Contact

For questions or support, feel free to open an issue in the repo.
