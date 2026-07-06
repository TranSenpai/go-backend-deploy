package gorm_client

import (
	"fmt"
	"go-backend/internal/common/env"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(env *env.Env) *gorm.DB {
	gormClient, err := gorm.Open(mysql.Open(env.DatabaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ [GORM] failed opening connection to mysql: %v", err)
	}

	err = gormClient.Raw("SELECT 1 + 1").Error
	if err != nil {
		log.Fatalf("❌ [GORM] failed connection to mysql: %v", err)
	}

	fmt.Println("✅ [GORM] Connection To MySQL Successfully")

	return gormClient
}
