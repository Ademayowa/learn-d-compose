# REST APIs With Go, PostgreSQL, and Gin</h2>

## <a name="tech-stack">⚙️ Tech Stack</a>

- Go
- PostgreSQL
- Gin

## <a name="quick-start">🤸 Quick Start</a>

Follow these steps to set up the project locally on your machine.

**Prerequisites**

Make sure you have the following installed on your machine:

- [Go](https://go.dev/doc/install)
- [Gin](https://gin-gonic.com/docs/quickstart/)
- [Docker](https://www.docker.com/)

**Cloning the Repository**

```bash
git clone https://github.com/Ademayowa/learn-d-compose.git
```

## Setup

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Create .env file

```bash
cp .env.example .env && echo "POSTGRES_PASSWORD=mysecurepassword123" >> .env
```

### 3. Running Docker compose to start all services

```bash
docker-compose up -d
```

### 3. Test the endpoint

```bash
curl -X POST http://localhost:8080/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Software Engineer",
    "description": "Develop software applications"
  }'
```

Open [http://localhost:8080/jobs](http://localhost:8080/jobs) in your browser to view all jobs.
