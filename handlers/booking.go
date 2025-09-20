package handlers

import (
	"app/db"
	"app/models"
	"app/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllBookings(c *fiber.Ctx) error {

	order, err := db.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BookingsResp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.BookingsResp{
		Status:  "success",
		Message: "All bookings fetched successfully",
		Data:    order,
	})
}

func GetBookingsById(c *fiber.Ctx) error {

	order_number := c.Params("id")

	order, err := db.GetOrderbyId(order_number)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BookingsResp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.BookingsResp{
		Status:  "success",
		Message: "Order with id '" + order_number + "' is found",
		Data:    order,
	})
}

func Book(c *fiber.Ctx) error {

	// Parse request body into EventBooking struct
	order := new(models.EventBooking)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BookingsResp{
			Status:  "error",
			Message: "Invalid request body: " + err.Error(),
		})
	}
	// Call the order service to register the booking
	status, err := services.Order(order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BookingsResp{
			Status:  "error",
			Message: err.Error(),
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(models.BookingsResp{
		Status:  "success",
		Message: "Booking successful",
		Data:    status,
	})
}

func DeleteBookingsById(c *fiber.Ctx) error {

	order_number := c.Params("id")

	_, err := db.DeleteOrder(order_number)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BookingsResp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.BookingsResp{
		Status:  "success",
		Message: "Order with id '" + order_number + "' has been deleted",
	})
}
