// Go connection Sample Code:
package db

import (
	"app/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Config variables
var (
	server   = "127.0.0.1"
	port     = 3306
	user     = "root"
	password = ""
	database = "books"
)

// Conn initializes and returns a MySQL database connection
func Conn() (*sql.DB, error) {
	// DSN format: username:password@tcp(host:port)/dbname?parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		user, password, server, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Set connection pool options
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Ping to verify connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Optionally store the connection globally
	DB = db

	return db, nil
}

func IdExists(booking_id string) (bool, error) {
	db, _ := Conn()

	const query = "SELECT 1 FROM bookings WHERE order_number = ?"
	row := db.QueryRow(query, booking_id)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Order(order *models.EventBooking) error {
	db, err1 := Conn()
	if err1 != nil {
		return err1
	}
	query := `INSERT INTO bookings (client_name, event_type, date, musician_type, order_number)
              VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, order.ClientName, order.EventType, order.Date, order.MusicianType, order.OrderNumber)

	return err
}

func DeleteOrder(order_number string) (string, error) {
	db, _ := Conn()

	// Start a new transaction
	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("could not begin transaction: %v", err)
	}

	// Defer a rollback in case of error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Delete order from the bookings table
	deleteUserQuery := "DELETE FROM bookings WHERE order_number = ?"
	_, err = tx.Exec(deleteUserQuery, order_number)
	if err != nil {
		return "", fmt.Errorf("could not delete order: %v", err)
	}

	// Commit the transaction if both deletions succeed
	if err = tx.Commit(); err != nil {
		return "", fmt.Errorf("could not commit transaction: %v", err)
	}

	return "ok", nil
}

func GetOrderbyId(order_number string) (models.EventBooking, error) {
	db, _ := Conn()
	var order models.EventBooking

	// Prepare the SQL query
	query := "SELECT client_name, event_type, date, musician_type FROM bookings WHERE order_number = ?"

	// Execute the query
	err := db.QueryRow(query, order_number).Scan(&order.ClientName, &order.EventType, &order.Date, &order.MusicianType)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.EventBooking{}, err
		}
		return models.EventBooking{}, err
	}

	return order, nil
}

func GetAllOrders() ([]models.EventBooking, error) {
	db, err := Conn()
	if err != nil {
		return nil, err
	}
	// Query to fetch data from the table
	rows, err := db.Query("SELECT client_name, event_type, date, musician_type, order_number FROM bookings")
	if err != nil {
		return nil, err
	}
	//defer rows.Close()

	// Slice to hold the results
	var orders []models.EventBooking

	// Iterate over the rows
	for rows.Next() {
		var order models.EventBooking
		err := rows.Scan(&order.ClientName, &order.EventType, &order.Date, &order.MusicianType, &order.OrderNumber)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// Print the results
	// for _, user := range users {
	// 	fmt.Printf("ID: %d, FName: %s, LName: %s, Email: %s\n", user.ID, user.FName, user.LName, user.Email)
	// }

	return orders, nil

}
