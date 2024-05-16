package main

import (
	"log/slog"
	"os"

	"github.com/AlissonBarbosa/shylockgo-api/models"
	"github.com/AlissonBarbosa/shylockgo-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
  l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  slog.SetDefault(l)

  models.ConnectDatabase()

  router := gin.Default()
  routes.RegisterRoutes(router)
  router.Run("0.0.0.0:5001")
}
