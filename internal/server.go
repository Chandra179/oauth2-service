package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Chandra179/auth-service/api"
	"github.com/Chandra179/auth-service/configs"
	"github.com/Chandra179/auth-service/pkg/encryptor"
	"github.com/Chandra179/auth-service/pkg/random"
	"github.com/Chandra179/auth-service/pkg/redis"
	"github.com/Chandra179/auth-service/pkg/serialization"
)

func StartServer() {
	// -------------
	// Configs
	// -------------
	config, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("err config", err)
	}
	// --------------
	// Logger
	// --------------
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// --------------
	// Redis
	// --------------
	rdb := redis.NewRedisClient("redis:6379", "", 0)
	// --------------
	// Serialization
	// --------------
	ser := serialization.NewGobSerialization()
	// --------------
	// Enryption
	// --------------
	aes, err := encryptor.NewAesEncryptor("0123456789abcdef") //16 bytes key
	if err != nil {
		fmt.Println("encryption err", err)
	}
	// --------------
	// Random
	// --------------
	rand := random.NewRandom(32)
	// --------------
	// Oauth
	// --------------
	googleOauth := NewGoogleOauth(config, rdb, logger, aes, ser, rand)
	// --------------
	// API setup
	// --------------
	api.SetupRoutes(googleOauth, config.GoogleOauth.RedirectURL)
	//---------------
	// Http Server
	// --------------
	log.Fatal(http.ListenAndServe(":8080", nil))
}
