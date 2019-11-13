// This file will have our database connection information, initialise our routes and start our servers

package controllers

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	"github.com/REST_API/api/models"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		Server.DB, err := gorm.Open("mysql", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error", err)
		}else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		Server.DB, err := gorm.Open("postgres","host=myhost port=myport user=gorm dbname=gorm password=mypassword")
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error", err)
		} else {
			fmt.Println("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	// server.DB.Debug().AutoMigrate will create model in passed database

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 9080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
	
}