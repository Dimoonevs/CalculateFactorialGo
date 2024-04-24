package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Dimoonevs/calculate/factorial/internal/handler"
	"github.com/Dimoonevs/calculate/factorial/internal/service"
	"github.com/spf13/viper"
)

func main() {
	// init configuration
	err := initConfig()
	if err != nil {
		log.Printf("error initializing config: %s", err.Error())
		os.Exit(1)
	}
	webPort := viper.GetString("port")

	// init service
	log.Printf("Starting Service on port %s", webPort)
	app := handler.NewAppFactorial(service.NewService())
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.NewRouter(),
	}

	// start server
	err = srv.ListenAndServe()
	if err != nil {
		log.Printf("Server stopped with error: %v", err)
		os.Exit(1)
	}

}

func initConfig() error {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
