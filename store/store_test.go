package store

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/A-pen-app/cache"
	"github.com/A-pen-app/kickstart/config"
	"github.com/A-pen-app/kickstart/database"
	"github.com/A-pen-app/kickstart/global"
	"github.com/A-pen-app/logging"
	"github.com/A-pen-app/tracing"
)

func TestMain(m *testing.M) {
	os.Setenv("TESTING", "true") // to inform different parts of the application that we are testing and perform accordingly

	// Create root context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Initializing resource for testing...")
	var projectID string = config.GetString("PROJECT_ID")
	if !config.GetBool("PRODUCTION_ENVIRONMENT") {
		projectID = ""
	}

	if err := logging.Initialize(&logging.Config{
		ProjectID:    projectID,
		Level:        logging.Level(config.GetUint("LOG_LEVEL")),
		Development:  !config.GetBool("PRODUCTION_ENVIRONMENT"),
		KeyRequestID: "request_id",
		KeyUserID:    "user_id",
		KeyError:     "err",
		KeyScope:     "scope",
	}); err != nil {
		panic(err)
	}
	defer logging.Finalize()

	// Setup tracing module
	env := "development"
	if config.GetBool("PRODUCTION_ENVIRONMENT") {
		env = "production"
	}
	tracing.Initialize(ctx, &tracing.Config{
		ProjectID:             config.GetString("PROJECT_ID"),
		TracerName:            "kickstart-api",
		ServiceName:           global.ServiceName,
		DeploymentEnvironment: env,
	})
	defer tracing.Finalize(ctx)

	cache.Initialize(&cache.Config{
		Type:     cache.TypeLocal,
		RedisURL: "localhost:6379",
		Prefix:   "local-dev",
	})
	defer cache.Finalize()

	// Setup database module.
	database.Initialize(ctx)
	defer database.Finalize()

	log.Println("Resource Initialized. Running tests...")
	exitVal := m.Run()

	os.Exit(exitVal)
}
