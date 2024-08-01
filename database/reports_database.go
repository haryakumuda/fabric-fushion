package database

import (
	"database/sql"
	"fmt"
	"strings"
)

func OrderReports(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
	select date(sales.order_date) as order_date,
		sum(sales_products.quantity) as total_quantity,
		sum(sales_products.quantity * products.price) as total_amount
	from sales
	join sales_products on sales.id = sales_products.sale_id
	join products on sales_products.product_id = products.id
	group by date(sales.order_date)
	order by date(sales.order_date) desc;
	`)
	if err != nil {
		return nil, fmt.Errorf("error fetching sales reports: %w", err)
	}
	defer rows.Close()

	var reports []map[string]interface{}
	for rows.Next() {
		var orderDate string
		var totalQuantity int
		var totalAmount float64

		if err := rows.Scan(&orderDate, &totalQuantity, &totalAmount); err != nil {
			return nil, fmt.Errorf("error scanning sales report: %w", err)
		}

		// Format the date
		dateParts := strings.Split(orderDate, "T")
		formattedDate := dateParts[0]

		report := map[string]interface{}{
			"order_date":     formattedDate,
			"total_quantity": totalQuantity,
			"total_amount":   totalAmount,
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func StockReports(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		select products.id, products.name, categories.category, products.stock
		from products
		join categories on products.category_id = categories.id
		order by products.name;
`)
	if err != nil {
		return nil, fmt.Errorf("error fetching stock reports: %w", err)
	}
	defer rows.Close()

	var reports []map[string]interface{}
	for rows.Next() {
		var productID, stock int
		var productName, categoryName string
		if err := rows.Scan(&productID, &productName, &categoryName, &stock); err != nil {
			return nil, fmt.Errorf("error scanning stock report: %w", err)
		}

		report := map[string]interface{}{
			"product_id":   productID,
			"product_name": productName,
			"category":     categoryName,
			"stock":        stock,
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func UserReports(db *sql.DB, customerID int) ([]map[string]interface{}, error) {
	query := `
	select 
		products.id as product_id,
		products.name as product_name,
		sales_products.quantity,
		(sales_products.quantity * products.price) as total_price,
		sales.order_date
	from sales
	join sales_products on sales.id = sales_products.sale_id
	join products on sales_products.product_id = products.id
	where sales.customer_id = ?
	order by sales.order_date desc;
	`

	rows, err := db.Query(query, customerID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user reports: %w", err)
	}
	defer rows.Close()

	var reports []map[string]interface{}
	for rows.Next() {
		var productID int
		var productName string
		var quantity int
		var totalPrice float64
		var orderDate string

		if err := rows.Scan(&productID, &productName, &quantity, &totalPrice, &orderDate); err != nil {
			return nil, fmt.Errorf("error scanning user report: %w", err)
		}

		//
		dateParts := strings.Split(orderDate, "T")
		formattedDate := dateParts[0]

		report := map[string]interface{}{
			"product_id":   productID,
			"product_name": productName,
			"quantity":     quantity,
			"total_price":  totalPrice,
			"order_date":   formattedDate,
		}
		reports = append(reports, report)
	}

	return reports, nil
}
