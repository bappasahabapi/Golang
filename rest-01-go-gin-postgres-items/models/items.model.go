package models

import (
	"go-gin-postgres-local/config"
)

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func GetAllItems() ([]Item, error) {
	rows, err := config.DB.Query("SELECT id, name, quantity FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func GetItemByID(id string) (Item, error) {
	var item Item
	err := config.DB.QueryRow("SELECT id, name, quantity FROM items WHERE id = $1", id).
		Scan(&item.ID, &item.Name, &item.Quantity)
	return item, err
}

func CreateItem(item *Item) error {
	return config.DB.QueryRow(
		"INSERT INTO items (name, quantity) VALUES ($1, $2) RETURNING id",
		item.Name, item.Quantity,
	).Scan(&item.ID)
}


func UpdateItem(id string, item *Item) error {
	_, err := config.DB.Exec("UPDATE items SET name = $1, quantity = $2 WHERE id = $3", item.Name, item.Quantity, id)
	return err
}

func DeleteItem(id string) error {
	_, err := config.DB.Exec("DELETE FROM items WHERE id = $1", id)
	return err
}
