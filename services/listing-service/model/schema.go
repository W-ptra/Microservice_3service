package model

import (
	"time"
)

type Listing struct {
	Id        	int       		`gorm:"primaryKey;autoIncrement" json:"id"` 
	UserId    	int				`json:"userId"`
	Price	  	int				`json:"price"`
	ListingType string			`json:"listingType"`
	CreatedAt 	time.Time     	`gorm:"autoCreateTime" json:"createdAt"` 
	UpdatedAt 	time.Time     	`gorm:"autoUpdateTime" json:"updatedAt"`
}