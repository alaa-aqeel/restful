package restful

import "github.com/gofiber/fiber"

type Controller struct{ 

}

func (ctrl *Controller) Get(ctx    *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }
func (ctrl *Controller) Show(ctx   *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }
func (ctrl *Controller) Post(ctx   *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }
func (ctrl *Controller) Update(ctx *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }
func (ctrl *Controller) Delete(ctx *fiber.Ctx){ ctx.Status(405).Send("Method Not Allowed") }