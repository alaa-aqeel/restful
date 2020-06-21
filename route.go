package restful 

import "github.com/gofiber/fiber"


func Resource(prefix string, group *fiber.Group, controller InterfaceResource) *fiber.Group {

	resource := group.Group(prefix)

	resource.Get("/", controller.Get)
	resource.Post("/", controller.Post)
	resource.Get("/:id", controller.Show)
	resource.Put("/:id", controller.Update)
	resource.Delete("/:id", controller.Delete)

	return resource
}


// Parser