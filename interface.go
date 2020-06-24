/*
	Package restful implements a simple library for restful-api with fiber.
	
*/
package restful 

import "github.com/gofiber/fiber"


// Interface for CRUD methods 
type InterfaceResource interface{

	Get(ctx *fiber.Ctx) 
	Show(ctx *fiber.Ctx) 
	Post(ctx *fiber.Ctx) 
	Delete(ctx *fiber.Ctx) 
	Update(ctx *fiber.Ctx) 
}
