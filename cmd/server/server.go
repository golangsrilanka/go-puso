package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	chiPrometheus "github.com/766b/chi-prometheus"
	"github.com/GolangSriLanka/go-puso/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/GolangSriLanka/go-puso/api/router"
	"github.com/GolangSriLanka/go-puso/config"
	_ "github.com/GolangSriLanka/go-puso/docs"
)

// @title Go Puso
// @version 0.0.1
// @description Golang Sri Lanka template repo

// @contact.name Golang Sri Lanka
// @contact.email golangsrilanka@mail.com

// @host golangsrilanka.github.io/go-puso
// @BasePath /api/v1

var RunServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start go-puso server",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Serve(ctx context.Context, r *chi.Mux) (err error) {
	port := config.GetEnv("server.PORT")
	if port == "" {
		port = "8080"
	}

	svc := http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
	}

	go func() {
		if err := svc.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	log.Println("Server running on port" + port)
	<-ctx.Done()
	log.Println("Server is starting to shutdown....")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := svc.Shutdown(ctxShutDown); err != nil {
		log.Println("Server was unable to shutdown")
	}

	log.Println("Server was shutdown successfully")

	return
}

func Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	port := config.GetEnv("server.PORT")

	r.Use(chiPrometheus.NewMiddleware("go-puso"))
	r.Handle("/metrics", promhttp.Handler())
	r.Mount("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))
	r.Mount("/healthz", router.HealthRoute())
	r.Mount("/api/v1", router.NewRouter(database.Database()).Route())

	go func() {
		oscall := <-quit
		log.Printf("oscall: %v\n", oscall)
		cancel()
	}()

	if err := Serve(ctx, r); err != nil {
		panic(err)
	}
}
