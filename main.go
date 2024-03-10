package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.println("Hello, World!")

	var router *gin.Engine = gin.Default()

	router.GET("/people", getPeople)
	router.POST("/people", postPeople)
	router.GET("/people/:id", getPersonByID)

	router.Run(":8000")
}

type person struct {
	ID string `json:"id"`
	NAME string `json:"name"`
}

var people = []person{
	{
		ID: "1",
		NAME: "Dhyan",
	},
	{
		ID: "2",
		NAME: "Madhav",
	},
	{
		ID: "3",
		NAME: "Shubh",
	},
}

func getPeople(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, people)
}

// add a person to people from JSON received in the request body
func postPeople(context *gin.Context) {
    var newPerson person

    // BindJSON to bind the received JSON to newPerson
    if err := context.BindJSON(&newPerson); err != nil {
        // log the error, respond and return
        fmt.Println(err)

        context.IndentedJSON(http.StatusBadRequest, gin.H{
            "message": "Invalid request",
        })

        return
    }

    // append the new person to people
    people = append(people, newPerson)

    // respond as IndentedJSON
    context.IndentedJSON(http.StatusCreated, newPerson)
}

func getPersonByID(context *gin.Context) {
    // get the id from request params
    var id string = context.Param("id")

    // Linear Search through people
    for _, p := range people {
        // respond and return if ID matched
        if p.ID == id {
            context.IndentedJSON(http.StatusOK, p)
            return
        }
    }

    // respond 404
    context.IndentedJSON(
        http.StatusNotFound,

        // refer https://pkg.go.dev/github.com/gin-gonic/gin#H
        gin.H{
            "message": "person not found",
        },
    )
}
