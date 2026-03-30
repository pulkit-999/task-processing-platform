# 🚀 Distributed Task Processing Platform (Go + Redis)

A scalable and fault-tolerant background job processing system built using **Go (Golang)** and **Redis**, designed to handle asynchronous workloads with high throughput and reliability.

---

## 🧠 Overview

This project implements a **distributed task processing platform** inspired by systems like Sidekiq and Celery.

It supports:

* Concurrent job execution using goroutines
* Retry with exponential backoff
* Dead Letter Queue (DLQ) for failed jobs
* Priority-based scheduling
* Dockerized deployment for easy setup

---

## ⚡ Key Features

* ⚙️ **Worker Pool (Concurrency)** — Multiple workers using goroutines for parallel job processing
* 🔁 **Retry Mechanism** — Automatic retries with exponential backoff
* ☠️ **Dead Letter Queue (DLQ)** — Failed jobs after max retries are stored separately
* 🎯 **Priority Queues** — High, Medium, Low priority scheduling
* 📊 **Metrics API** — Track processed, failed, and queued jobs
* 🐳 **Dockerized Setup** — Run entire system using Docker Compose

---

## 🏗️ Architecture

```
Client
   │
   ▼
API Server (Gin)
   │
   ▼
Redis (Queue)
   │
   ▼
Worker Pool (Goroutines)
   │
   ├── Retry (Exponential Backoff)
   ├── Priority Scheduling
   └── Dead Letter Queue (DLQ)
   │
   ▼
Metrics API
```

---

## 📁 Project Structure

```
cmd/
  ├── api/        # REST API (producer)
  └── worker/     # Worker service (consumer)

internal/
  ├── models/     # Job definitions
  ├── queue/      # Redis queue logic
  ├── worker/     # Worker pool implementation
  ├── metrics/    # Metrics tracking
```

---

## ⚙️ Tech Stack

* **Go (Golang)**
* **Gin (HTTP framework)**
* **Redis (Message Queue)**
* **Docker & Docker Compose**

---

## ▶️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/pulkit-999/task-processing-platform.git
cd task-processing-platform
```

---

### 2. Run using Docker

```bash
docker-compose up --build
```

---

### 3. API will be available at:

```
http://localhost:8080
```

---

## 📬 API Endpoints

### ➤ Create Job

**POST /job**

```json
{
  "type": "email",
  "payload": "Send welcome email",
  "priority": "high"
}
```

---

### ➤ Get Metrics

**GET /metrics**

```json
{
  "processed": 10,
  "failed": 2,
  "queued": 5
}
```

---

## 🧪 How It Works

1. Client sends job to API
2. API pushes job to Redis queue
3. Worker pool consumes jobs concurrently
4. Failed jobs are retried with exponential backoff
5. After max retries → moved to Dead Letter Queue
6. Metrics API tracks system performance

---

## 🔍 Redis Operations (Internal)

* `LPUSH` → Add job to queue
* `BRPOP` → Worker consumes job
* `DLQ` → Stores failed jobs

---

## 🎯 Use Cases

* Background email processing
* Asynchronous task execution
* Microservices job handling
* Event-driven systems

---

## 📌 Future Improvements

* Redis Streams support
* Rate limiting
* Monitoring dashboard
* Kubernetes deployment

---

## 👨‍💻 Author

**Pulkit Agrawal**

* GitHub: https://github.com/pulkit-999
* LinkedIn: https://linkedin.com/in/pulkit-agrawal14

---

## ⭐ If you found this useful, give it a star!
