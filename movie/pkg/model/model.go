package model

import "movieexample.com/metadata/pkg/model"

// Movie Details includes movie metadata and its aggregate rating
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata`
}
