package main



func main()  {
	
}

func SetupRoutes(app *fiber.App){
	app.Get("/url", routes.Resolve)

}