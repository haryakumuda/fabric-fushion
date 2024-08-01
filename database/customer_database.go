package database

import (
	"database/sql"
	"encoding/json"
	"fabric-fushion/helper"
	"fabric-fushion/model"
	"log"
	"time"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) InsertSale(CustomerID int, products map[int]int) error {
	productsJSON, err := json.Marshal(helper.ConvertToJSON(products))
	if err != nil {
		return err
	}

	// call store procedure
	query := `Call InsertSale(?,?,?)`
	_, err = db.DB.Exec(query, CustomerID, time.Now(), productsJSON)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) ShowProducts() (*sql.Rows, error) {
	rows, err := db.DB.Query("CALL GetProductsWithCategory()")
	if err != nil {
		log.Fatal(err)
	}
	return rows, nil
}

func (db *Database) GetOrderHistory(CustomerID int) ([]model.Sales, error) {
	var orders []model.Sales

	salesQuery := `SELECT id, order_date FROM sales WHERE customer_id = ?`
	rows, err := db.DB.Query(salesQuery, CustomerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Sales
		if err := rows.Scan(&order.ID, &order.OrderDate); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (db *Database) GetProductsForOrder(saleID int64) (*sql.Rows, error) {
	query := `
		SELECT
			sp.id,
			sp.sale_id,
			sp.product_id,
			sp.quantity,
			p.name,
			p.price,
			p.category_id,
			c.category
		FROM sales_products sp
		JOIN products p ON sp.product_id = p.id
		JOIN categories c ON p.category_id = c.id
		WHERE sp.sale_id = ?`
	rows, err := db.DB.Query(query, saleID)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// UpdateProductStock update stok
func (db *Database) UpdateProductStock(productID int64, quantity int) error {
	query := `UPDATE products SET stock = stock + ? WHERE id = ?`

	_, err := db.DB.Exec(query, quantity, productID)
	if err != nil {
		return err
	}

	return nil
}
