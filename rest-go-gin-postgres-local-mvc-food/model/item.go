package models

import (
	"database/sql"
	"fmt"
	"log"

	"go-gin-postgres-local/utils" // Adjust module path if different
)

// Item struct to map to our database table
type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity"`
}

// GetAllItems retrieves all items from the database
func GetAllItems(db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT id, name, quantity FROM items")
	if err != nil {
		log.Printf(utils.ColorRed+"Error querying items: %v"+utils.ColorReset, err)
		return nil, fmt.Errorf("failed to retrieve items")
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
			log.Printf(utils.ColorRed+"Error scanning item row: %v"+utils.ColorReset, err)
			return nil, fmt.Errorf("failed to process item data")
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		log.Printf(utils.ColorRed+"Error after iterating rows: %v"+utils.ColorReset, err)
		return nil, fmt.Errorf("error processing item results")
	}
	return items, nil
}

// GetItemByID retrieves a single item by its ID from the database
func GetItemByID(db *sql.DB, id string) (*Item, error) {
	var item Item
	row := db.QueryRow("SELECT id, name, quantity FROM items WHERE id = $1", id)
	err := row.Scan(&item.ID, &item.Name, &item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Item not found, not necessarily an application error
		}
		log.Printf(utils.ColorRed+"Error querying single item with ID %s: %v"+utils.ColorReset, id, err)
		return nil, fmt.Errorf("failed to retrieve item: %w", err)
	}
	return &item, nil
}

// CreateItem inserts a new item into the database and updates its ID
func CreateItem(db *sql.DB, item *Item) error {
	err := db.QueryRow(
		"INSERT INTO items (name, quantity) VALUES ($1, $2) RETURNING id",
		item.Name, item.Quantity,
	).Scan(&item.ID)

	if err != nil {
		log.Printf(utils.ColorRed+"Error creating item: %v"+utils.ColorReset, err)
		return fmt.Errorf("failed to create item: %w", err)
	}
	return nil
}

// UpdateItem updates an existing item in the database
// It returns the updated item or an error if the item is not found or another error occurs.
func UpdateItem(db *sql.DB, id string, itemData *Item) (*Item, error) {
	var updatedItem Item
	err := db.QueryRow(
		"UPDATE items SET name = $1, quantity = $2 WHERE id = $3 RETURNING id, name, quantity",
		itemData.Name, itemData.Quantity, id,
	).Scan(&updatedItem.ID, &updatedItem.Name, &updatedItem.Quantity)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item with ID %s not found for update", id)
		}
		log.Printf(utils.ColorRed+"Error updating item with ID %s: %v"+utils.ColorReset, id, err)
		return nil, fmt.Errorf("failed to update item: %w", err)
	}
	return &updatedItem, nil
}

// DeleteItem removes an item from the database by its ID
// It returns an error if the item is not found or another error occurs.
func DeleteItem(db *sql.DB, id string) error {
	result, err := db.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		log.Printf(utils.ColorRed+"Error deleting item with ID %s: %v"+utils.ColorReset, id, err)
		return fmt.Errorf("failed to delete item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf(utils.ColorRed+"Error getting rows affected after deleting item with ID %s: %v"+utils.ColorReset, id, err)
		return fmt.Errorf("failed to confirm deletion of item: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("item with ID %s not found for deletion", id)
	}
	return nil
}