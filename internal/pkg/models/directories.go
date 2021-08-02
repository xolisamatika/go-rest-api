package models

import "time"

type Directory struct {
	FullPath         string    `json:"fullPath"`
	Size             int64     `json:"size"`
	IsDirective      bool      `json:"isDirective"`
	ModificationTime time.Time `json:"modifcationTime"`
}
