package main

import (
	"log"
	"net/http"

	gojwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/julioc98/shion/cmd/api/handler"
	"github.com/julioc98/shion/cmd/api/middleware"
	"github.com/julioc98/shion/cmd/api/router"
	"github.com/julioc98/shion/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/julioc98/shion/pkg/entity"
	"github.com/julioc98/shion/pkg/env"
	"github.com/julioc98/shion/pkg/guardian"
	"github.com/julioc98/shion/pkg/password"
	"github.com/julioc98/shion/pkg/repository"
	"github.com/julioc98/shion/pkg/usecase/user"
	_ "github.com/shaj13/libcache/fifo"

	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
)

func handlerHi(w http.ResponseWriter, r *http.Request) {
	msg := "Ola, Seja bem vindo ao Shion!!"
	log.Println(msg)
	w.Write([]byte(msg))
}

var strategy union.Union
var keeper jwt.SecretsKeeper

func main() {

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := provideDBConn(providePostgresDialector(conf), provideGormConfig())
	if err != nil {
		log.Fatal(err)
	}

	conn.AutoMigrate(&entity.User{})

	// go-guardian
	keeper = jwt.StaticSecret{
		ID:     "secret-id",
		Secret: []byte("secret"),
		Method: gojwt.SigningMethodHS256,
	}

	validate := validator.New()

	r := mux.NewRouter()
	r.Use(middleware.Logging)

	pwd := password.NewService()

	userRep := repository.NewUserRepository(conn)

	userService := user.NewService(userRep, pwd)

	authMethod := guardian.New(userService.ValidateEmailAndPassword, keeper)

	userHandler := handler.NewUserHTTPHandler(userService, validate, authMethod)

	r.Use(authMethod.AuthMiddleware)

	router.SetUserRoutes(userHandler, r.PathPrefix("/users").Subrouter())

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	port := env.Get("PORT", "5001")
	log.Printf(`%s listening on port: %s `, env.Get("APP", "shion"), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func provideDBConn(dialector gorm.Dialector, gConf *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(dialector, gConf)
}

func provideGormConfig() *gorm.Config {
	return &gorm.Config{}
}

func providePostgresDialector(conf *config.Configuration) gorm.Dialector {
	return postgres.Open(conf.Database.URL)
}
