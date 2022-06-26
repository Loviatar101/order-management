package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/loviatar101/order-consumer/database"
	"github.com/loviatar101/order-consumer/routes"
)


func welcome(c *fiber.Ctx)error{
	return c.SendString("welcome to api")
}

func setupRoutes(app *fiber.App){
//welcome endpoint
app.Get("/api",welcome)
//create user
app.Post("/api/users",routes.CreateUser)
//get users
app.Get("/api/users",routes.GetUsers)
//get by id 
app.Get("/api/users/:id",routes.GetUser)
//update user
app.Put("/api/users/:id",routes.UpdateUser)
//delete user
app.Delete("/api/users/:id",routes.DeleteUser)



//products
//create product

app.Post("/api/products",routes.CreateProduct)
//GetProducts
app.Get("/api/products",routes.GetProducts)
//get by id 
app.Get("/api/products/:id",routes.GetProduct)
//update products
app.Put("/api/products/:id",routes.UpdateProduct)
//delete products
app.Delete("/api/products/:id",routes.DeleteProduct)


//create order
app.Post("/api/order",routes.CreateOrder)
//get orders
app.Get("/api/order",routes.GetOrders)
//get order by id
app.Get("/api/order/:id",routes.GetOrder)


}

func main() {

	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)	
	app.Listen(":3000")
}