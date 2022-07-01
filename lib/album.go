package lib

import "fmt"

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (a Album) String() string {
	return fmt.Sprintf("%s: %s", a.Title, a.Artist)
}
