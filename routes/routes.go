package routes

import (
	"app/handlers"

	"github.com/gofiber/fiber/v2"
)

// The function that runs all routes and starts the server
func Run() {
	app := fiber.New()

	app.Get("/bookings/getall", handlers.GetAllBookings)
	app.Get("/bookings/get/:id", handlers.GetBookingsById)
	app.Delete("/bookings/delete/:id", handlers.DeleteBookingsById)
	app.Post("/bookings", handlers.Book)

	app.Listen(":2020")
}
