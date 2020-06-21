# Restful-fiber :-)

##### Make resources route `CRUD` with fiber 
##### Methods Get Post Show Delete Update 

##### Install 
	go get github.com/alaaProg/restful

##### Controller 	
	
	import "github.com/alaaProg/restful"

	....
	....

	type UserCtrl struct{
		restful.Controller 
	}

	func (ctrl *UserCtrl) Get(ctx *fiber.Ctx){

		ctx.Send("Get User")
	}

	....
	....

#### Route 	

	import "github.com/alaaProg/restful"

	...
	...

	api := app.Group("/api/") 

	restful.Resource("/user", api, new(controller.UserCtrl))  // return *fiber.Group 

	.... 
	....