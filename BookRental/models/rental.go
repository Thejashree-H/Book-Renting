package models

type Rental struct {
	ID 		 int
	UserID 	 int
	BookID 	 int
	IsRented bool
}