# Stock Board Go

A CLI application that fetches and displays real-time stock prices using the Finnhub API. It uses Goroutines to concurrently fetch data for multiple stock symbols and updates the console every 10 seconds.

## 🚀 Features

- **Real-time Data**: Fetches live stock prices from the Finnhub API.
- **Concurrency**: Uses Goroutines and WaitGroups to fetch multiple stock prices concurrently.
- **Auto Refresh**: Automatically updates and displays prices every 10 seconds.
- **Clean Architecture**: Separates concerns into Controllers, Services, Repositories, and Models.

## 🛠️ Tech Stack

- **Go 1.25.1**
- **godotenv** (for environment variable management)

## 📂 Project Structure

```text
stock-board-go/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   └── stock/
│       ├── controllers/
│       │   └── stock.controller.go # Handles requests from main and calls services
│       ├── models/
│       │   └── stock.go            # Data structures
│       ├── repositories/
│       │   └── stock.repository.go # Data access layer (currently maps data to models)
│       └── services/
│           └── stock.service.go    # Business logic and external API calls (Finnhub)
├── .env                            # Environment variables
├── go.mod
└── README.md
```

## ⚙️ Setup and Installation

1. **Clone Repository**

   ```bash
   git clone <repository-url>
   cd stock-board-go
   ```

2. **Initialize Go Modules**

   ```bash
   go mod tidy
   ```

3. **Environment Variables**

   Create a `.env` file in the root directory and provide your Finnhub API key and desired symbols:

   ```env
   # Configuration for API
   API_URL=https://finnhub.io/api/v1/quote
   SYMBOL=AAPL,MSFT,GOOGL
   
   # Authentication 
   FINNHUB_API_KEY="your_finnhub_api_key_here"
   ```

## 🏃‍♂️ Running the Application

You can run the application directly using:

```bash
go run cmd/server/main.go
```

The application will start fetching data and print it to the console. It will continue to refresh every 10 seconds until you terminate it (e.g., using `Ctrl+C`).

**Example Output:**

```text
Live Stock Price:
Ticker: AAPL, Price: 173.50, Last Update: 2026-04-26 14:10:00
Ticker: GOOGL, Price: 140.20, Last Update: 2026-04-26 14:10:00
Ticker: MSFT, Price: 410.00, Last Update: 2026-04-26 14:10:00
------------------------------------------------
```

## 🏗️ Architecture

This application follows a **Clean Architecture** pattern adapted for a CLI tool, divided into 4 main layers:

1. **Controller Layer** (`internal/stock/controllers`)
   - Acts as the entry point from the `main` function.
   - Delegates requests to the Service layer.

2. **Service Layer** (`internal/stock/services`)
   - Contains the core business logic.
   - Responsible for making HTTP requests to the external Finnhub API.
   - Coordinates with the Repository layer.

3. **Repository Layer** (`internal/stock/repositories`)
   - Abstracts data creation and storage.
   - Currently implements a mock storage that maps the fetched price to the `Stock` model.

4. **Model Layer** (`internal/stock/models`)
   - Defines the core data structures (`structs`) used throughout the application.

## 📝 License

MIT
