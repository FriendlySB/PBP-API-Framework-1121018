package main

import (
	"Tugas-Explorasi-1-PBP-Framework-API/controller"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/cors"
)

func main() {
	//Test Framework Martini + Golangci
	//1 = admin
	//2 = user
	m := martini.Classic()

	//CORS global
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"localhost:8492"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	//CORS untuk sebuah route/access point saja
	allowCORSHandler := cors.Allow(&cors.Options{
		AllowOrigins:     []string{"localhost:8492"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	//Grouped routers
	m.Group("/users", func(r martini.Router) {
		r.Post("", controller.Authenticate(controller.InsertUser, 1))
		r.Get("", controller.Authenticate(controller.GetAllUsers, 1))
		r.Put("/:user_id", controller.UpdateUser)
		r.Delete("/:user_id", controller.DeleteUser)
	})
	//CORS khusus untuk fungsi ini
	m.Delete("/userscors/:user_id", allowCORSHandler, controller.DeleteUser)

	//Router biasa
	m.Post("/login", controller.LoginUser)
	m.Post("/logout", controller.LogoutUser)

	//Ganti port
	m.RunOnAddr(":8492")

}
