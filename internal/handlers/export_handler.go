package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/mallacharmi/polyglot-export-engine/internal/models"
	"github.com/mallacharmi/polyglot-export-engine/internal/services"
)

type ExportHandler struct {
	service *services.ExportService
}

func NewExportHandler(service *services.ExportService) *ExportHandler {
	return &ExportHandler{service: service}
}

func (h *ExportHandler) CreateExport(c *gin.Context) {

	var req models.CreateExportRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	job, err := h.service.CreateExport(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"exportId": job.ID,
		"status":   job.Status,
	})
}

func (h *ExportHandler) GetExport(c *gin.Context) {

	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	job, exists := h.service.GetExport(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "export not found"})
		return
	}

	c.JSON(http.StatusOK, job)
}