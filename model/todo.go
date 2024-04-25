package model

import "time"

type (
	// A TODO expresses ...
	TODO struct {
		ID          int       `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	// タイムスタンプはtime.Timeらしい

	// A CreateTODORequest expresses ...
	CreateTODORequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}

	// A CreateTODOResponse expresses ...
	CreateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	// A ReadTODORequest expresses ...
	ReadTODORequest struct {
		ID          int    `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}

	// A ReadTODOResponse expresses ...
	ReadTODOResponse struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}

	// A UpdateTODORequest expresses ...
	UpdateTODORequest struct{}
	// A UpdateTODOResponse expresses ...
	UpdateTODOResponse struct{}

	// A DeleteTODORequest expresses ...
	DeleteTODORequest struct{}
	// A DeleteTODOResponse expresses ...
	DeleteTODOResponse struct{}
)
