// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	Title string `json:"title"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	Address1   *string `json:"address1"`
	Address2   *string `json:"address2"`
	Address3   *string `json:"address3"`
	City       *string `json:"city"`
	Prefecture *string `json:"prefecture"`
	PostalCode *string `json:"postalCode"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Restaurant struct {
	Name        string       `json:"name"`
	ID          string       `json:"id"`
	URL         *string      `json:"url"`
	Phone       *string      `json:"phone"`
	Price       *string      `json:"price"`
	Categories  []*Category  `json:"categories"`
	Location    *Location    `json:"location"`
	Coordinates *Coordinates `json:"coordinates"`
	Photos      []string     `json:"photos"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}