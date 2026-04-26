package repositories

import (
	"database/sql"
	"fmt"
	"stock-board-go/internal/stock/models"
)

type StockRepository interface {
	SaveStock(stock *models.Stock) error
}

type stockRepository struct {
	db *sql.DB
}

func NewStockRepository(db *sql.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) SaveStock(stock *models.Stock) error {
	query := `
		INSERT INTO stock_prices (ticker, price, last_update)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(query, stock.Ticker, stock.Price, stock.LastUpdate)
	if err != nil {
		return fmt.Errorf("failed to save stock: %w", err)
	}
	return nil
}
