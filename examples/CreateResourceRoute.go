package main 

import "github.com/gofiber/fiber"
import "github.com/alaaProg/restful"


type UserCtrl {
	restful.Controller 
}


func (ctrl *UserCtrl) Get(ctx *fiber.Ctx){
	ctx.JSON(fiber.Map{
			"msg": "Hello, I am Get Method!"
		})
}

func (ctrl *UserCtrl) Post(ctx *fiber.Ctx){

	ctx.JSON(fiber.Map{
			"msg": "Hello, I am Post Method!"
		})
}


func main(){
	// create app
	app := fiber.New()

	// create group api 
	api := app.Group("/api/")
		
	// create Create and read method 
	// other method  Update and Delete returen method not allow
	restful.Resource("/user", api, new(UserCtrl), func(ctx fiber.Ctx){
			// middleware for Resource
			ctx.Next()
		})




	app.Listen(":8080")

}