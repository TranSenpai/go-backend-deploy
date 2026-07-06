package env

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	IsProduction bool
	Port         string
	Host         string
	DatabaseUrl  string

	ExpiresAtAccessToken time.Duration
	SecretAccessToken    string

	ExpiresAtRefreshToken time.Duration
	SecretRefreshToken    string

	GoogleClientId     string
	GoogleClientSecret string
	GoogleRedirectUrl  string

	DomainFe string
	DomainBe string

	CloudinaryUrl string

	RedisAddr string
	RedisPass string

	ElasticAddrs           string
	ElasticUser            string
	ElasticPassword        string
	ElasticCertFingerprint string

	RabbitMQURL string
}

func New() *Env {
	godotenv.Load()

	isProduction := os.Getenv("IS_PRODUCTION") == "true"
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	databaseUrl := os.Getenv("DATABASE_URL")

	// ACCESS TOKEN
	expiresAtAccessTokenString := os.Getenv("EXPIRES_AT_ACCESS_TOKEN")
	expiresAtAccessToken := getDuration(expiresAtAccessTokenString)
	secretAccessToken := os.Getenv("SECRET_ACCESS_TOKEN")

	// REFRESH TOKEN
	expiresAtRefreshTokenString := os.Getenv("EXPIRES_AT_REFRESH_TOKEN")
	expiresAtRefreshToken := getDuration(expiresAtRefreshTokenString)
	secretRefreshToken := os.Getenv("SECRET_REFRESH_TOKEN")

	// GOOGLE
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirectUrl := os.Getenv("GOOGLE_REDIRECT_URL")

	domainFe := os.Getenv("DOMAIN_FE")
	domainBe := os.Getenv("DOMAIN_BE")

	cloudinaryUrl := os.Getenv("CLOUDINARY_URL")

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")

	elasticAddrs := os.Getenv("ELASTIC_ADDRS")
	elasticUser := os.Getenv("ELASTIC_USER")
	elasticPassword := os.Getenv("ELASTIC_PASSWORD")
	elasticCertFingerprint := os.Getenv("ELASTIC_CERT_FINGERPRINT")

	rabbitMQURL := os.Getenv("RABBIT_MQ_URL")

	fmt.Println("isProduction", isProduction)
	fmt.Println("port", port)
	fmt.Println("host", host)
	fmt.Println("databaseUrl", databaseUrl)

	fmt.Println("expiresAtAccessTokenString", expiresAtAccessTokenString)
	fmt.Println("secretAccessToken", secretAccessToken)

	fmt.Println("expiresAtRefreshTokenString", expiresAtRefreshTokenString)
	fmt.Println("secretRefreshToken", secretRefreshToken)

	fmt.Println("googleClientId", googleClientId)
	fmt.Println("googleClientSecret", googleClientSecret)
	fmt.Println("googleRedirectUrl", googleRedirectUrl)

	fmt.Println("domainFe", domainFe)
	fmt.Println("domainBe", domainBe)

	fmt.Println("cloudinaryUrl", cloudinaryUrl)

	fmt.Println("redisAddr", redisAddr)
	fmt.Println("redisPass", redisPass)

	fmt.Println("elasticAddrs", elasticAddrs)
	fmt.Println("elasticUser", elasticUser)
	fmt.Println("elasticPassword", elasticPassword)
	fmt.Println("elasticCertFingerprint", elasticCertFingerprint)

	return &Env{
		IsProduction:           isProduction,
		Port:                   port,
		Host:                   host,
		DatabaseUrl:            databaseUrl,
		ExpiresAtAccessToken:   expiresAtAccessToken,
		SecretAccessToken:      secretAccessToken,
		ExpiresAtRefreshToken:  expiresAtRefreshToken,
		SecretRefreshToken:     secretRefreshToken,
		GoogleClientId:         googleClientId,
		GoogleClientSecret:     googleClientSecret,
		GoogleRedirectUrl:      googleRedirectUrl,
		DomainFe:               domainFe,
		CloudinaryUrl:          cloudinaryUrl,
		DomainBe:               domainBe,
		RedisAddr:              redisAddr,
		RedisPass:              redisPass,
		ElasticAddrs:           elasticAddrs,
		ElasticUser:            elasticUser,
		ElasticPassword:        elasticPassword,
		ElasticCertFingerprint: elasticCertFingerprint,
		RabbitMQURL:            rabbitMQURL,
	}
}

func getDuration(durationString string) time.Duration {
	durationTime, err := time.ParseDuration(durationString)
	if err != nil {
		log.Fatal("Parser expiresAtAccessToken error")
	}
	return durationTime
}
