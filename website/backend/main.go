package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kmchuc/angellist-take-home/tree/main/website/backend/util"
	"log"
)

func main() {
	fmt.Println("Hello World")

	app := fiber.New()

	opportunity := util.Allocate{}

	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	app.Post("/api/invest", func(ctx *fiber.Ctx) error {
		request := &util.Allocate{}

		if pErr := ctx.BodyParser(request); pErr != nil {
			return fmt.Errorf("failed to parse request body: %w", pErr)
		}

		opportunity = *request

		prorates, pErr := util.ProrateCalculator(opportunity.AllocationAmount, opportunity.InvestorAmounts)
		if pErr != nil {
			return fmt.Errorf("error calculating prorate: %w", pErr)
		}

		return ctx.JSON(prorates)
	})

	log.Fatal(app.Listen(":4000"))
}
