package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"example.com/graphql-blog-api/graph"
	"example.com/graphql-blog-api/graph/resolver"
	"example.com/graphql-blog-api/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func RunServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := godotenv.Load(".env");
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connexion à la base de données
	Database := graph.Connect()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{DB: Database}}))
	
	// Configuration des en-têtes CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Autoriser les requêtes depuis ce domaine
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", corsHandler(srv))

	log.Print("🚀 Server running on http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
