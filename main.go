package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kevinfinalboss/FinOps/database"
	"github.com/kevinfinalboss/FinOps/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro carregando arquivo .env")
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	asciiArt := figure.NewFigure("FinOps", "", true)
	asciiArt.Print()

	router := gin.Default()

	routes.RegisterRoutes(router)

	port := getPort()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	database.ConnectToMongoDB()
	defer database.DisconnectFromMongoDB()
	
	go func() {
		log.Printf("Servidor rodando na porta: %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Falha ao iniciar o servidor: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Desligando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Servidor forçado a desligar: ", err)
	}

	log.Println("Servidor desligado.")
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
