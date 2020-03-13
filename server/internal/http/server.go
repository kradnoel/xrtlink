package http

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

var port string

func init() {
	godotenv.Load(".env")
}

type server struct{ Port string }

func New() *server {
	port := os.Getenv("PORT")
	if port == "" {
		log.Panic("PORT is not set!!!")
		return nil
	}
	return &server{Port: port}
}

func (s *server) Run() {
	//r := gin.Default()
	//r.GET("/seed", SeedLinks)
	/*r.GET("/:id", getLink)
	r.POST("/", postLink)
	r.OPTIONS("/", optionsLink)
	r.Run(":" + s.Port)*/

	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"localhost:3000"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	//contentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

	router.Use(cors.Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//router.Use(middleware.AllowContentType([]string{"application/x-www-form-urlencoded"}))
	//router.Use(middleware.AllowContentType("application/x-www-form-urlencoded"))
	router.Get("/{id}", getLink)
	router.Post("/", postLink)

	http.ListenAndServe(":"+s.Port, router)
}

func Execute() {
	server := New()
	server.Run()
}
