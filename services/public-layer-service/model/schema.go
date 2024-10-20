package model

import "time"

type ErrorRespond struct{
	Result bool 	`json:"respond"`
	Message string 	`json:"message"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRespond struct{
	Result bool		`json:"result"`
	User  User 		`json:"user"`
}

type UsersRespond struct{
	Result bool			`json:"result"`
	User  []User 		`json:"users"`
}

type Listing struct {
	Id        	int       		`json:"id"` 
	UserId    	int				`json:"userId"`
	Price	  	int				`json:"price"`
	ListingType string			`json:"listingType"`
	CreatedAt 	time.Time     	`json:"createdAt"` 
	UpdatedAt 	time.Time     	`json:"updatedAt"`
}

type ListingRespond struct{
	Result bool			`json:"result"`
	Listing []Listing 	`json:"listing"`
}

type ListingCreateRespond struct{
	Result bool			`json:"result"`
	Listing Listing 	`json:"listing"`
}

type ListingWithUser struct {
	Id        	int       		`json:"id"` 
	Price	  	int				`json:"price"`
	ListingType string			`json:"listingType"`
	User 		User			`json:"user"`
	CreatedAt 	time.Time     	`json:"createdAt"` 
	UpdatedAt 	time.Time     	`json:"updatedAt"`
}

type ListingWithUserRespond struct{
	Result bool					`json:"result"`
	Listing  []ListingWithUser 	`json:"listing"`
}
