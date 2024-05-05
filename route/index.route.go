package route

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Get("/products", handler.ProductHandlerGetAll)
	r.Post("/products", handler.ProductHandlerCreate)
	r.Get("/products/:id", handler.ProductHandlerGetById)
	r.Put("/products/:id", handler.ProducthandlerUpdate)
	r.Delete("/products/:id", handler.ProductHandlerDelete)

}
