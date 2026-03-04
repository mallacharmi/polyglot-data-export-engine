package models

import "github.com/google/uuid"

type ColumnMapping struct {
	Source string `json:"source" binding:"required"`
	Target string `json:"target" binding:"required"`
}

type CreateExportRequest struct {
	Format      string          `json:"format" binding:"required"`
	Columns     []ColumnMapping `json:"columns" binding:"required"`
	Compression string          `json:"compression,omitempty"`
}

type ExportJob struct {
	ID          uuid.UUID
	Format      string
	Columns     []ColumnMapping
	Compression string
	Status      string
}