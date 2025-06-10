package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings" // For checking "not found" error messages

	models "go-gin-postgres-local/model"
	"go-gin-postgres-local/utils" // Adjust module path if different

	"github.com/gin-gonic/gin"
)

// ItemHandler holds the database connection
type ItemHandler struct {
	DB *sql.DB
}

// NewItemHandler creates a new ItemHandler
func NewItemHandler(db *sql.DB) *ItemHandler {
	return &ItemHandler{DB: db}
}

// GetItems handles the GET /items request
func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := models.GetAllItems(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetItem handles the GET /items/:id request
func (h *ItemHandler) GetItem(c *gin.Context) {
	id := c.Param("id")
	item, err := models.GetItemByID(h.DB, id)
	if err != nil {
		// The model's GetItemByID already logs the specific DB error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve item"})
		return
	}
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// CreateItem handles the POST /items request
func (h *ItemHandler) CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.CreateItem(h.DB, &newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseBytes, marshalErr := json.MarshalIndent(newItem, "", "  ")
	if marshalErr != nil {
		log.Printf(utils.ColorGreen+"Successfully created item: %+v"+utils.ColorReset, newItem)
	} else {
		log.Printf(utils.ColorGreen+"Successfully created item. Response:\n%s"+utils.ColorReset, string(responseBytes))
	}
	c.JSON(http.StatusCreated, newItem)
}

// UpdateItem handles the PUT /items/:id request
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var itemData models.Item

	if err := c.ShouldBindJSON(&itemData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedItem, err := models.UpdateItem(h.DB, id, &itemData)
	if err != nil {
		if strings.Contains(err.Error(), "not found for update") {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Item with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	responseBytes, marshalErr := json.MarshalIndent(updatedItem, "", "  ")
	if marshalErr != nil {
		log.Printf(utils.ColorGreen+"Successfully updated item with ID %s: %+v"+utils.ColorReset, id, updatedItem)
	} else {
		log.Printf(utils.ColorGreen+"Successfully updated item with ID %s. Response:\n%s"+utils.ColorReset, id, string(responseBytes))
	}
	c.JSON(http.StatusOK, updatedItem)
}

// DeleteItem handles the DELETE /items/:id request
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	err := models.DeleteItem(h.DB, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found for deletion") {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Item with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	log.Printf(utils.ColorGreen+"Successfully deleted item with ID %s."+utils.ColorReset, id)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Item with ID %s successfully deleted", id)})
}