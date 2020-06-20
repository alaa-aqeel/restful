# restful
##### Make resources route `CRUD` with fiber 
##### Methods Get Post Show Delete Update 

##### Install 
	go get github.com/alaaProg/restful

##### Controller 	
	
	import "github.com/alaaProg/restful"

	type User struct{
		restful.Controller
	}

	func (user *User) Get(ctx *fiber.Ctx){

		ctx.Send("Get User")
	}

#### Route 

	api := app.Group("/api/") // fiber Group 

	restful.Resource("/user", api, new(ctrl.User)) 

