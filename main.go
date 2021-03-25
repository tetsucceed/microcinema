package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var Addr = ":8080"

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

func timeMustParse(year string) time.Time {
	tm, err := time.Parse("2006", year)
	if err != nil {
		panic(err)
	}
	return tm
}

func movieListHandler(rc *gin.Context) {
	var mt MovieId
	err := rc.ShouldBindUri(&mt)
	if err != nil {
		fmt.Println("oops")
	}

	rc.Set("Content-Type", "application/json; charset=utf-8")
	mm := []Movie{
		Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg",
			"https://youtu.be/qtRKdVHc-cE", true, timeMustParse("1999"),
			"triller"},
		Movie{1, "Крестный отец", "/static/posters/father.jpg",
			"https://youtu.be/ar1SHxgeZUc", false, timeMustParse("1988"),
			"drama"},
		Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg",
			"https://youtu.be/s7EdQ4FqbhY", true, timeMustParse("1996"),
			"comedy"},
	}

	if mt == (MovieId{}) {
		rc.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "response": mm})
		return
	} else {
		id, err := strconv.Atoi(mt.Id)
		if err != nil {
			rc.JSON(http.StatusBadRequest, gin.H{"error": "Somthing bad happen"})
		}
		if id > -1 && id < 3 {
			rc.JSON(http.StatusOK, mm[id])
		}
	}

}

func main() {
	r := gin.Default()
	group := r.Group("/api")
	{
		group.GET("/movies", movieListHandler)
		group.GET("/movies/:id", movieListHandler)
	}
	_ = r.Run(Addr)
}
