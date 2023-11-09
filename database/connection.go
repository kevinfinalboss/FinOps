package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

func ConnectToMongoDB() {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("A variável de ambiente MONGO_URL não está definida")
	}

	clientOptions := options.Client().ApplyURI(mongoURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoDBClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	err = MongoDBClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Falha ao pingar o MongoDB: %v", err)
	}

	log.Println("Conexão com MongoDB estabelecida com sucesso.")
}

func DisconnectFromMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := MongoDBClient.Disconnect(ctx); err != nil {
		log.Fatalf("Falha ao desconectar do MongoDB: %v", err)
	}

	log.Println("Desconectado do MongoDB com sucesso.")
}
