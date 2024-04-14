package dashboard1

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"../models"
)

type Dashboard1Handler struct {
	db *sql.DB
}

func NewDashboard1Handler(db *sql.DB) *Dashboard1Handler {
	return &Dashboard1Handler{db: db}
}

func (h *Dashboard1Handler) HandleDashboard1(c *gin.Context) {

	data, err := h.fetchData(parameters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "dashboard1.html", gin.H{"data": string(jsonData)})
}

func (h *Dashboard1Handler) fetchData(parameters map[string]string) ([]models.Data, error) {

	sql := `SELECT id, name, company, value, description
           FROM your_table
           WHERE ... conditions based on parameters ...;`

	rows, err := h.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Data
	for rows.Next() {
		var d models.Data
		err := rows.Scan(&d.ID, &d.Name, &d.Company, &d.Value, &d.Description)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	return data, nil
}
