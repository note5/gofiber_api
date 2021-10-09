package services

import (
	"github.com/gofiber/fiber/v2"

	"fmt"
)

//add custom middleware

func CheckMiddleware(c *fiber.Ctx) error {
	fmt.Print("======== MIDDLEWARE RUNS======", string(c.Body()))
	c.Next()
	return nil
}
func RestrictTo (c *fiber.Ctx) error {
	fmt.Print("======== RESTRICT TO MIDDLEWARE ======", string(c.Body()))
	c.Next()
	return nil
}