package models

type Book struct {
	ID 			  int
	Title 		  string
	Author		  string
	Description   string
	PublishedYear int
	Available     bool
}