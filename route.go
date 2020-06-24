/*
	Package restful implements a simple library for restful-api with fiber.
	
*/
package restful 

import "github.com/gofiber/fiber"


/*
	Resource assigns the typical "CRUD" routes to a controller with a single line of code.

	Parameters
		prefiex string: prefiex route group 
		group *fiber.Group : Group route
		controller InterfaceResource: Controller struct
		handlers ...fiber.Handler: Middleware funcation 


	return group for resource
*/
func Resource(prefix string, group *fiber.Group, 
		controller InterfaceResource, handlers ...fiber.Handler) *fiber.Group {

	resource := group.Group(prefix)
	
	for _, handler := range handlers {
		resource.Use(handler)
	}

	resource.Get("/", controller.Get)
	resource.Post("/", controller.Post)
	resource.Get("/:id", controller.Show)
	resource.Put("/:id", controller.Update)
	resource.Delete("/:id", controller.Delete)

	return resource
}
