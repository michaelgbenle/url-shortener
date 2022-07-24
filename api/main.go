package main



func main()  {
	
}

func SetupRoutes(app *fiber.App){
	app.get("/url", routes.Resolve)

}