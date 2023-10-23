package server

import (
	"context"
	"net/http"
	"time"

	chiPrometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"

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
		fx.New(
			fx.Supply(cmd),
			Init,
			fx.Invoke(Run),
		).Run()
	},
}

func Run(lc fx.Lifecycle, ro *router.Router) {
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
	if port == "" {
		port = "8080"
	}

	r.Use(chiPrometheus.NewMiddleware("go-puso"))
	r.Handle("/metrics", promhttp.Handler())
	r.Mount("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))
	r.Mount("/healthz", router.HealthRoute())
	r.Mount("/api/v1", ro.Route())

	svc := http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Server running on port: " + port)

				err := svc.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server is starting to shutdown....")

			return svc.Shutdown(ctx)
		},
	})
}
