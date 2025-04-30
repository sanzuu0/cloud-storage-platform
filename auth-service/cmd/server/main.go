package main

import (
	"github.com/sanzuu0/cloud-storage-platform/auth-service/config"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/app"
	"log"
)

// TODO: Точка входа в приложение
//  - Загрузить конфигурацию (config)
//  - Инициализировать зависимости (Postgres, Redis, Kafka)
//  - Создать экземпляр Service
//  - Настроить маршруты (HTTP Server)
//  - Запустить HTTP сервер

func main() {
	cfg := config.Load()

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
