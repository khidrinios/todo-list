package cmd

import (
	"context"
	"khidr/todo/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func API() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Start go-todo in api mode.",
		Run: func(cmd *cobra.Command, args []string) {
			dsn := "host=localhost user=khidr password=khidr dbname=onboarding port=5432 sslmode=disable"
			log.Println("Starting go-cassiopeia api...")
			runAPI(dsn)
		},
	}
}

func runAPI(dsn string) {
	engine := router.InitAPI(dsn)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
}
