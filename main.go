package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"petstore-api/config"
	"petstore-api/handlers"
	"petstore-api/repositories"
	"petstore-api/routes"
	"petstore-api/services"

	_ "petstore-api/docs" // Swagger docs
)

// @title Pet Store API
// @version 1.0
// @description A pet store management API with sellers and pets
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

func main() {
	fmt.Println("Starting Pet Store API...")

	fmt.Println("Initializing database...")
	db := config.InitDB()
	fmt.Println("Database initialized successfully")

	sellerRepo := repositories.NewSellerRepository(db)
	petRepo := repositories.NewPetRepository(db)

	sellerService := services.NewSellerService(sellerRepo, petRepo)
	petService := services.NewPetService(petRepo, sellerRepo)

	sellerHandler := handlers.NewSellerHandler(sellerService)
	petHandler := handlers.NewPetHandler(petService)

	router := routes.SetupRoutes(sellerHandler, petHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		fmt.Println("🚀 Server starting on http://localhost:8080")
		fmt.Println("📚 Swagger UI available at: http://localhost:8080/swagger/index.html")
		fmt.Println("\nAvailable endpoints:")
		fmt.Println("  GET    /sellers")
		fmt.Println("  POST   /sellers")
		fmt.Println("  GET    /sellers/{id}")
		fmt.Println("  PUT    /sellers/{id}")
		fmt.Println("  DELETE /sellers/{id}")
		fmt.Println("  GET    /pets")
		fmt.Println("  POST   /pets")
		fmt.Println("  GET    /pets/{id}")
		fmt.Println("  PUT    /pets/{id}")
		fmt.Println("  DELETE /pets/{id}")
		fmt.Println("\nPress Ctrl+C to stop the server")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	fmt.Println("✅ Server stopped gracefully")
}
