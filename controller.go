/*
	Package restful implements a simple library for restful-api with fiber.
	
*/
package restful

import "github.com/gofiber/fiber"

// 
type Messages map[string]map[string]string 
//
type Map map[string]string 


// 
type Controller struct{
	MessageValid Messages
	
}

 
// Get / 
func (ctrl *Controller) Get(ctx    *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }

// Get  /:id
func (ctrl *Controller) Show(ctx   *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }

// Post /
func (ctrl *Controller) Post(ctx   *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }

// Put /:id
func (ctrl *Controller) Update(ctx *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }

// Delete /:id 
func (ctrl *Controller) Delete(ctx *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }