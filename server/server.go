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

const defaultPort = "8080"

func main() {
	databaseName := helpers.GoDotEnvVariable("DBNAME")
	databaseUser := helpers.GoDotEnvVariable("DBUSER")
	databasePassword := helpers.GoDotEnvVariable("DBPASSWORD")

	port := helpers.GoDotEnvVariable("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB(databaseUser, databasePassword, databaseName)

	/*
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	*/

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())

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
