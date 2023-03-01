package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kmchuc/angellist-take-home/tree/main/website/backend/_pkg/util"
	. "github.com/tbxark/g4vercel"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "POST",
	}))

	opportunity := util.Allocate{}

	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	app.Post("/api/invest", func(ctx *fiber.Ctx) error {
		request := new(util.Allocate)

		if pErr := ctx.BodyParser(request); pErr != nil {
			return fmt.Errorf("failed to parse request body: %w", pErr)
		}

		opportunity = *request

		prorates, pErr := util.AllocationCalculator(opportunity.AllocationAmount, opportunity.InvestorAmounts)
		if pErr != nil {
			return fmt.Errorf("error calculating prorate: %w", pErr)
		}

		return ctx.JSON(prorates)
	})

	log.Fatal(app.Listen(":4000"))
}
