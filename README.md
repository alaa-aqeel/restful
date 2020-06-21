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

##### Validation  

	type UserModel struct{

		Name string `required:"true" match:"^[a-z]$" length:2,4`
	}

	func (ctrl *UserCtrl) Get(ctx *fiber.Ctx){

		user := new(models.UserModel)
		ctx.BodyParser(user)

		// Create custom messages 
		ctrl.MessageValid = restful.Messages{
			"name" : restful.Map{ // required, match, length 
				"required": "%s field is required" // return `name filed ....`
				"match": "%s field is match" // return `name filed ....`
				"length": "%s field is between %s" // return `name filed between 2,4`
			}
		}

		if messages, err := ctrl.Valid(user); err {

			ctx.Status(401).JSON(messages)
		}
		ctx.Send("Get User")
	}

	....
	....
