package main

import (
	"ChatApp/models"
	"ChatApp/routers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("ERROR : ", err)
	}
	r := routers.SetupRouter()

	port := os.Getenv("port")

	// For run on requested port
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //localhost
	}

	dbCon := models.GetGormDB()
	models.AutoMigrate(dbCon)

	logrus.Info("<<<<<<< Server Started >>>>>>>")
	err = r.Run(":" + port)
	if err != nil {
		logrus.Error("ERROR : ", err)
	}

}