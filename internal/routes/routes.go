package routes

import (
	"workout_tracker/internal/handlers"
	"workout_tracker/internal/middleware"
	"workout_tracker/internal/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", services.Signup)
	r.Post("/login", services.Login)
}

func WebRoutes(app *fiber.App) {

	app.Static("/js", "./web/js")
	app.Static("/", "./web")

	app.Get("/index", middleware.RequireAuth, func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./web/index.html")
	})
	app.Get("/signup", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./web/signup.html")
	})

	app.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./web/login.html")
	})

	app.Post("/newset", handlers.NewSet)
	app.Get("/user/sets", handlers.LoadSets)
	app.Post("/user/updateset", handlers.UpdateSet)
	app.Post("/user/deleteset", handlers.DeleteSet)
	app.Get("/user/updateform", handlers.UpdateForm)

}
