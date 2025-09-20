# Book Project 📚  

This is a Golang-based project that provides a structured backend for handling bookings, authentication, cards, and related services. The project is modularized into routes, handlers, models, services, utilities, and database connection management.

---

## 📂 Project Structure  

```
Book/
│── go.mod                 # Go module file
│── go.sum                 # Go dependencies checksum
│── main.go                # Application entry point
│
├── db/
│   └── conn.go             # Database connection logic
│
├── handlers/
│   ├── .env                # Environment variables file
│   └── cards.go            # Request handlers for cards
│
├── models/
│   └── auth_struct.go      # Structs for authentication
│
├── routes/
│   └── routes.go           # Route definitions
│
├── services/
│   └── booking.go          # Booking-related business logic
│
├── utils/
│   └── helper.go           # Helper functions
│
└── .idea/                  # IDE (JetBrains/GoLand) project settings
```

---

## 🚀 Getting Started  

### 1. Clone the repository  
```bash
git clone https://github.com/yourusername/book.git
cd book
```

### 2. Install dependencies  
```bash
go mod tidy
```

### 3. Set environment variables  
Create a `.env` file in the `handlers/` directory or root project directory with required environment variables (such as DB credentials, API keys, etc.). Example:  
```
DB_HOST=localhost
DB_USER=root
DB_PASS=yourpassword
DB_NAME=bookdb
PORT=8080
```

### 4. Run the application  
```bash
go run main.go
```

---

## ⚡ Features  

- **Database Connection** – Centralized in `db/conn.go`.  
- **Card Handlers** – API endpoints for managing cards (`handlers/cards.go`).  
- **Authentication Models** – Structs to define auth payloads (`models/auth_struct.go`).  
- **Booking Services** – Handles business logic for bookings (`services/booking.go`).  
- **Routing** – Managed in `routes/routes.go`.  
- **Utility Functions** – Common helpers in `utils/helper.go`.  

---

## 🛠️ Tech Stack  

- **Language**: Go (Golang)  
- **Database**: MySQL / PostgreSQL (based on `.env` config)  
- **Architecture**: RESTful service with modularized packages  

---

## 📌 TODOs  

- Add JWT authentication middleware  
- Improve error handling with structured responses  
- Write unit tests for services and handlers  
- Add Dockerfile for containerized deployment  
