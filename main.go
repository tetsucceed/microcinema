package main

import (
	"context"
	"dbo"
	"fmt"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"strconv"
	"time"
)

var Addr = ":8080"

func timeMustParse(year string) time.Time {
	tm, err := time.Parse("2006", year)
	if err != nil {
		panic(err)
	}
	return tm
}

func movieListHandler(rc *gin.Context) {
	var mt model.MovieId
	err := rc.ShouldBindUri(&mt)
	if err != nil {
		fmt.Println("oops")
	}

	rc.Set("Content-Type", "application/json; charset=utf-8")
	mm := dbo.LoadMovies()

	if mt == (model.MovieId{}) {
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
	dbo.GetConnectionDb()
	defer dbo.DB.Close(context.Background())
	dbo.ApplyMigration(dbo.DBUrl)

	r := gin.Default()
	group := r.Group("/api")
	{
		group.GET("/movies", movieListHandler)
		group.GET("/movies/:id", movieListHandler)
	}
	_ = r.Run(Addr)
}
