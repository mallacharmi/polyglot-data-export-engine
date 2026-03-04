package services

import (
	"errors"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/mallacharmi/polyglot-export-engine/internal/models"
)

var allowedFormats = map[string]bool{
	"csv":     true,
	"json":    true,
	"xml":     true,
	"parquet": true,
}

var allowedColumns = map[string]bool{
	"id":         true,
	"created_at": true,
	"name":       true,
	"value":      true,
	"metadata":   true,
}

type ExportService struct {
	store map[uuid.UUID]models.ExportJob
	mutex sync.RWMutex
}

func NewExportService() *ExportService {
	return &ExportService{
		store: make(map[uuid.UUID]models.ExportJob),
	}
}

func (s *ExportService) CreateExport(req models.CreateExportRequest) (models.ExportJob, error) {

	format := strings.ToLower(req.Format)

	if !allowedFormats[format] {
		return models.ExportJob{}, errors.New("invalid format")
	}

	if req.Compression != "" && req.Compression != "gzip" {
		return models.ExportJob{}, errors.New("invalid compression type")
	}

	for _, col := range req.Columns {
		if !allowedColumns[col.Source] {
			return models.ExportJob{}, errors.New("invalid column: " + col.Source)
		}
	}

	job := models.ExportJob{
		ID:          uuid.New(),
		Format:      format,
		Columns:     req.Columns,
		Compression: req.Compression,
		Status:      "pending",
	}

	s.mutex.Lock()
	s.store[job.ID] = job
	s.mutex.Unlock()

	return job, nil
}

func (s *ExportService) GetExport(id uuid.UUID) (models.ExportJob, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	job, exists := s.store[id]
	return job, exists
}