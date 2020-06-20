package restful 

import "github.com/gofiber/fiber"

type InterfaceResource interface{

	Get(ctx *fiber.Ctx) 
	Show(ctx *fiber.Ctx) 
	Post(ctx *fiber.Ctx) 
	Delete(ctx *fiber.Ctx) 
	Update(ctx *fiber.Ctx) 
}
