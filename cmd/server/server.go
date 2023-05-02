package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/andreis3/graphql-goexpert/graph"
	"github.com/andreis3/graphql-goexpert/internal/database"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryService: categoryDB,
		CourseService:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
