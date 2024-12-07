package main

import (
	"gorm.io/gorm"
)

type Feature struct {
	gorm.Model
	Title         string `json:"title"`
	Description   string `json:"description"`
	CustomerScore int    `json:"customer_score"`
	InternalScore int    `json:"internal_score"`
	SalesScore    int    `json:"sales_score"`
	ImpactScore   int    `json:"impact_score"`
	Type          string `json:"type"`
}
