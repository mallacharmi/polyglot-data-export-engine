# Polyglot Data Export Engine

A Dockerized Go-based backend service that allows exporting data from a PostgreSQL database using customizable column mappings and export formats.

This project demonstrates backend microservice development using **Go (Golang)**, **Gin Web Framework**, **PostgreSQL**, and **Docker**.

---

## Features

- REST API built with Gin framework
- Export data from PostgreSQL tables
- Column mapping support
- CSV export format
- Dockerized microservice
- Health monitoring endpoint
- Clean project architecture (services, handlers, models)

---

## Tech Stack

- Go (Golang)
- Gin Web Framework
- PostgreSQL
- Docker & Docker Compose

---

## Project Structure

```
polyglot-data-export-engine
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── database/
│   ├── handlers/
│   ├── models/
│   ├── services/
│   └── writers/
│
├── seeds/
│   └── init-db.sh
│
├── tests/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## Prerequisites

Make sure the following are installed:

- Docker
- Docker Compose
- Git
- Go (optional if running locally)

---

## Running the Project

Clone the repository:

```
git clone https://github.com/yourusername/polyglot-data-export-engine.git
```

Navigate into the project folder:

```
cd polyglot-data-export-engine
```

Start the application using Docker:

```
docker-compose up --build
```

The application will start at:

```
http://localhost:8080
```

---

## API Endpoints

### Health Check

```
GET /health
```

Example:

```
http://localhost:8080/health
```

Response:

```
{
  "status": "healthy"
}
```

---

### Create Export Job

```
POST /exports
```

Example Request:

```
curl -X POST http://localhost:8080/exports \
-H "Content-Type: application/json" \
-d '{
  "format": "csv",
  "columns": [
    {"source":"id","target":"id"},
    {"source":"name","target":"name"},
    {"source":"value","target":"value"}
  ]
}'
```

Example Response:

```
{
  "exportId": "416dca00-4f40-45b2-b1bf-558d4c8bb6ad",
  "status": "pending"
}
```

---

### Get Export Status

```
GET /exports/{id}
```

Example:

```
curl http://localhost:8080/exports/416dca00-4f40-45b2-b1bf-558d4c8bb6ad
```

Example Response:

```
{
  "ID": "416dca00-4f40-45b2-b1bf-558d4c8bb6ad",
  "Format": "csv",
  "Status": "pending"
}
```

---

## Database

The project initializes PostgreSQL automatically using the seed script:

```
seeds/init-db.sh
```

Table created:

```
records
```

Columns:

- id
- created_at
- name
- value
- metadata

---

## Running Containers

Check running containers:

```
docker ps
```

Stop containers:

```
docker-compose down
```

---

## Author

Malla Charmi
BTech Student – Aditya University

---

## License

This project is for educational purposes.
