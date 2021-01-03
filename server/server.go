package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/vuegqltodo/server/database"
	"github.com/hyperxpizza/vuegqltodo/server/graph"
	"github.com/hyperxpizza/vuegqltodo/server/graph/generated"
	"github.com/hyperxpizza/vuegqltodo/server/helpers"
)

const defaultPort = "8888"

func main() {
	databaseUser := "hyperxpizza"
	databasePassword := "secret_password"
	databaseName := "vuegqltodo"

	port := helpers.GoDotEnvVariable("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB(databaseUser, databasePassword, databaseName)

	router := gin.Default()

	//router.Use(CORS())
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:8080"},
		AllowMethods:  []string{"OPTIONS", "POST", "GET", "PUT"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/query", graphqlHandler())
	//router.GET("/", playgroundHandler())

	router.Run(":" + port)
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
