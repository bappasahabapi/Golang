package handlers

import (
	"database/sql"
	"encoding/json" // Import for JSON marshaling (for pretty printing)
	models "go-gin-postgres-local/model"
	"go-gin-postgres-local/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ItemHandler holds the database connection (or service layer)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	err := models.CreateItem(h.DB, &newItem) // newItem is updated with ID by this function
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// --- Log the created item (which now includes the ID) ---
	// For pretty printing the JSON in the log:
	responseBytes, marshalErr := json.MarshalIndent(newItem, "", "  ") // Marshal with indentation
	if marshalErr != nil {
		// Fallback to simple logging if marshaling fails (should be rare)
		log.Printf(utils.ColorGreen+"Successfully created item: %+v"+utils.ColorReset, newItem)
	} else {
		log.Printf(utils.ColorGreen+"Successfully created item. Response:\n%s"+utils.ColorReset, string(responseBytes))
	}
	// --- End of logging ---

	c.JSON(http.StatusCreated, newItem)
}


