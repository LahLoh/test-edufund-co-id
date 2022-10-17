package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/shandysiswandi/test-edufund-co-id/handler"
	"github.com/shandysiswandi/test-edufund-co-id/pkg/clock"
	"github.com/shandysiswandi/test-edufund-co-id/pkg/security"
	"github.com/shandysiswandi/test-edufund-co-id/pkg/token"
	"github.com/shandysiswandi/test-edufund-co-id/repository"
	"github.com/shandysiswandi/test-edufund-co-id/service"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")

	repoOpt := repository.MysqlOption{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		Options:  os.Getenv("DB_OPTIONS"),
	}

	repo, err := repository.New(repoOpt)
	if err != nil {
		log.Fatalln("failed to open db", err)
	}

	if err = repo.Ping(); err != nil {
		log.Fatalln("failed to connect db", err)
	}

	if err = repo.Migrate("./migration"); err != nil {
		log.Fatalln("failed to migrate db", err)
	}

	hasher := security.Bcrypt{Cost: 14}
	token := token.JWT{Secret: []byte(os.Getenv("JWT_SECRET"))}
	clocker := clock.Time{}

	svc := service.New(repo, &hasher, &token, &clocker)
	hand := handler.New(svc)

	r := mux.NewRouter()
	r.Use(handler.JsonMiddleware)
	r.HandleFunc("/v1/register", hand.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", hand.Login).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":50501",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("server running on port 50501")
	log.Fatal(srv.ListenAndServe())
}
