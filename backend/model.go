package main

import (
	"gorm.io/gorm"
)

type Feature struct {
	gorm.Model
	Title         string `json:"title"`
	Description   string `json:"description"`
	CustomerScore int    `json:"customer_score"`
	DevScore      int    `json:"dev_score"`
	SalesScore    int    `json:"sales_score"`
	Type          string `json:"type"`
	Estimate      string `json:"estimate"`
}

type Milestone struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
