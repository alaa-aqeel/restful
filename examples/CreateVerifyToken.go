package main 

import "os"
import "fmt"
import "github.com/gofiber/fiber"
import "github.com/alaaProg/restful"


type UserCtrl {
	restful.Controller 
}


var Username string 

type FormValid struct{

	Username 	 string `required:"true" `
	Passwrod 	 string `required:"true"`
}

// Login
func (ctrl *UserCtrl) Login(ctx *fiber.Ctx){
	form := new(FormValid)
	ctx.BodyParser(form)

	if msgs, err := ctrl.Valid(form); err != nil {
		ctx.Status(401).JSON(msgs)
		return
	}

	token, err := ctrl.CreateToken(form.Username, 10) // 10m
	if err != nil {
		ctx.Status(401).JSON(err.Error())
	}


	ctx.JSON(fiber.Map{
			"token": token
		})
}

func (ctrl *UserCtrl) Home(ctx *fiber.Ctx){


	ctx.JSON(fiber.Map{
			"token": fmt.Sprintf("Hello %s", Username)
		})
}

func AuthMidle(ctx *fiber.Ctx){
	token := ctx.Get("Authorization")

	if token == ""{
		ctx.Status(401).JSON(fiber.Map{
			"msg": "Unauthorized",
		})
		return 
	}

	token = strings.Split(token, " ")[1]

	tkn, user, err := restful.VerifyToken(token)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{
			"msg": err.Error(),
		})
		return 
	}

	if !tkn.Valid{
		ctx.Status(401).JSON(fiber.Map{
			"msg": "Unauthorized",
		})
		return 
	}


	Username = user.Id
	ctx.Next()
}

func main(){
	// create app
	app := fiber.New()

	os.Setenv("TOEKN_SECRET_KEY", "TEST_SECRET_KEY")

	// create group api 
	api := app.Group("/api/")
		
	api.Post("/login", new(UserCtrl).Login)
	api.Use(AuthMidle)
	api.Post("/user",  new(UserCtrl).Home)

	app.Listen(":8080")
}