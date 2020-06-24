package main 

import "github.com/gofiber/fiber"
import "github.com/alaaProg/restful"


type UserCtrl {
	restful.Controller 
}


type FormValid struct{

	name 	 string `required:"true" match:"^[a-z]+$"`
	age 	 int    `required:"true"`
	message  string `required:"true" length:"100,1000"`

}

func (ctrl *UserCtrl) Post(ctx *fiber.Ctx){
	form := new(FormValid)
	ctx.BodyParser(form)

	if msgs, err := ctrl.Valid(form); err != nil {
		ctx.Status(401).JSON(msgs)
		return
	}

	ctx.JSON(fiber.Map{
			"msg": "Hello, I am Post Method!"
		})
}


func main(){
	// create app
	app := fiber.New()

	// create group api 
	api := app.Group("/api/")
		
	// create CRUP
	restful.Resource("/user", api, new(UserCtrl), func(ctx fiber.Ctx){
			// middleware for Resource
			ctx.Next()
		})


	app.Listen(":8080")
}