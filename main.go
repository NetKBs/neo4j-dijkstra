package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"example.com/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driver *config.DB
var ctx context.Context

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	ctx = context.Background()
	url := os.Getenv("NEO4J_URL")
	user := os.Getenv("NEO4J_USER")
	pass := os.Getenv("NEO4J_PASS")

	driver = config.NewDB(url, user, pass)
	err := driver.Connect(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established")
	defer driver.Close(ctx)

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/path", GetShortestPath)

	r.Run()
}

type Objetive struct {
	Start string `json:"start" binding:"required"`
	End   string `json:"end" binding:"required"`
}

type Path struct {
	Path []string `json:"path"`
}

func GetShortestPath(c *gin.Context) {
	var obj Objetive
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := neo4j.ExecuteQuery(ctx, driver.GetDB(),
		`MATCH (start:Estacion {nombre: $name1}), (end:Estacion {nombre: $name2})
		CALL apoc.algo.dijkstra(
			start,
			end,
			"CONECTA_CON>",
			"tiempo"
		) YIELD path
		RETURN CASE WHEN path IS NOT NULL THEN [n in nodes(path) | n.nombre] ELSE [] END AS estaciones`,

		map[string]any{
			"name1": obj.Start,
			"name2": obj.End,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(result.Records) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontro la ruta"})
	}

	stations, _ := result.Records[0].Get("estaciones")
	stationsSlice := stations.([]any)

	path := make([]string, len(stationsSlice))
	for i, station := range stationsSlice {
		path[i] = station.(string)
	}

	response := Path{Path: path}

	c.JSON(http.StatusOK, response)

}
