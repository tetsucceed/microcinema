package model

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Poster      string    `json:"poster"`
	MovieUrl    string    `json:"movie_url"`
	IsPaid      bool      `json:"is_paid"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       string    `json:"genre"`
}

type MovieId struct {
	Id string `uri:"id" binding:"required"`
}
