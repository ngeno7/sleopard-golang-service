package repositories

import (
	"database/sql"
	"log"

	"leopard.test/v2/internal/db"
	"leopard.test/v2/internal/models"
)


func CreateCustomer(customer models.Customer) (*int64, error) {

	conn, err := db.Connect()
	if err != nil {
		// http.Error(w, "DB connection failed", http.StatusInternalServerError)
		return nil, nil
	}
	defer conn.Close()

	query := "INSERT INTO customers (first_name, last_name, phone, location, preferred_product) VALUES (?, ?, ?, ?, ?)"
	result, err := conn.Exec(query, customer.FirstName, customer.LastName, customer.Phone, customer.Location, customer.PreferredProduct)
	if err != nil {
		return nil, nil
	}

	id, _ := result.LastInsertId()
	
	return &id, nil
}

func GetCustomer(id int64) (*models.Customer, error) {

	

	conn, err := db.Connect()
	if err != nil {

		return nil, err
	}
	defer conn.Close()

	var customer models.Customer
	query := "SELECT id, first_name, last_name, phone, location FROM customers WHERE id = ? LIMIT 1"

	err = conn.QueryRow(query, id).Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Phone, &customer.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			
			return nil, err
		}

		return nil, err
	}

	return &customer, nil
}

func GetCustomers()(*[]models.Customer, error) {

	conn, err := db.Connect()
	if err != nil {

		return nil, err
	}
	defer conn.Close()

	var customers []models.Customer

	row,err := conn.Query("SELECT id, first_name, last_name, location, preferred_product FROM customers")

	if err != nil {
		log.Printf("Querying customers :%v", err)
		return nil, err
	}
	for row.Next() {
		var customer models.Customer

		if err := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Location, &customer.PreferredProduct); err != nil {
            return nil, err
        }
		customers = append(customers, customer)
	}

	return &customers, nil
}