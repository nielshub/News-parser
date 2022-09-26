package main

import (
	"context"
	"incrowd/src/internal/handlers"
	"incrowd/src/internal/repositories"
	"incrowd/src/internal/services"
	"incrowd/src/log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Start logger and load environment variables
	log.Init("debug")
	//err := godotenv.Load("/usr/local/bin/variables.env")
	err := godotenv.Load("../../env/variables.env")
	if err != nil {
		log.Logger.Error().Msgf("Variables file not found... Error: %s", err)
		panic(err)
	}
	log.Logger.Info().Msgf("Environment variables loaded")

	// Connect to mongoDB in Docker
	dbURL := os.Getenv("DBURL")
	sportNewsCollectionName := os.Getenv("SPORTNEWSCOLLECTIONNAME")
	dataBaseName := os.Getenv("DATABASENAME")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Logger.Error().Msgf("DB client has been disconnected. Error: %s", err)
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Logger.Error().Msgf("Error not connected to mongoDB. Error: %s", err)
		return
	}

	log.Logger.Info().Msgf("Connected to users DB")

	// Create repositories / services / handlers and app
	db := client.Database(dataBaseName).Collection(sportNewsCollectionName)
	NonRelationalSportNewsDBRepository := repositories.NewMongoDBRepository(sportNewsCollectionName, db)
	sportNewsService := services.NewSportNewsService(NonRelationalSportNewsDBRepository)
	cronPullService := services.NewCronPullService(NonRelationalSportNewsDBRepository)

	r := gin.Default()
	app := r.Group("/")

	handlers.NewHealthHandler(app)
	handlers.NewSportNewsHandler(app, sportNewsService)

	//cronPullService.CronPullNewsRoutine(context.Background())

	//Launch go cron routine
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Minutes().Do(func() { cronPullService.CronPullNewsRoutine(context.Background()) })

	// Run server
	log.Logger.Info().Msgf("Starting server")
	err = r.Run(":8080")
	if err != nil {
		log.Logger.Error().Msgf("Error running the server on port 8080. Error: %s", err)
	}

	log.Logger.Info().Msgf("Stopping server")

}
