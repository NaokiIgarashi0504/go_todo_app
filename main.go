package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
	/*
			fmt.Println(config.Config.Port)
			fmt.Println(config.Config.SQLDriver)
			fmt.Println(config.Config.DbName)
			fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	/*
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)
		u.CreateUser()
	*/

	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo ")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	/*
		user, _ := models.GetUser(3)
		user.CreateTodo("Third Todo")
	*/

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		t, _ := models.GetTodo(3)
		t.DeleteTodo()
	*/
}
