package main

import (
	"fmt"
	"log"
	"os"
	"stock-board-go/internal/stock/controllers"
	"stock-board-go/internal/stock/models"
	"stock-board-go/internal/stock/repositories"
	"stock-board-go/internal/stock/services"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	stockRepo := repositories.NewStockRepository()
	stockService := services.NewStockService(stockRepo)
	stockController := controllers.NewStockController(stockService)

	symbol := os.Getenv("SYMBOL")
	fmt.Println("Live Stock Price:")
	for {
		stockChannel := make(chan models.Stock)
		var wg sync.WaitGroup

		for _, symbol := range strings.Split(symbol, ",") {
			wg.Add(1)
			go stockController.GetStock(symbol, &wg, stockChannel)
		}

		go func() {
			wg.Wait()
			close(stockChannel)
		}()

		for stock := range stockChannel {
			fmt.Printf("Ticker: %s, Price: %.2f, Last Update: %s\n", stock.Ticker, stock.Price, stock.LastUpdate.Format("2006-01-02 15:04:05"))
		}

		fmt.Println("------------------------------------------------")
		time.Sleep(10 * time.Second)
	}
}
