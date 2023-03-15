package main

import (
	"Tugas-Explorasi-1-PBP-Framework-API/controller"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()

	m.Group("/users", func(r martini.Router) {
		r.Get("", controller.Authenticate(controller.GetAllUsers, 1))
		r.Post("", controller.Authenticate(controller.InsertUser, 1))
		r.Put("", controller.UpdateUser)
		r.Delete("/:user_id", controller.DeleteUser)
	})

	m.Post("/login", controller.LoginUser)
	m.Post("/logout", controller.LogoutUser)

	m.RunOnAddr(":8492")

}
