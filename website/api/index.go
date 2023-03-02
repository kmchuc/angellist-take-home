package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kmchuc/angellist-take-home/tree/main/website/backend/_pkg/util"
	. "github.com/tbxark/g4vercel"
	"net/http"
)

//Handler is the default handler
func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))

	var allocation util.Allocate

	server.POST("/api/invest", func(c *Context) {
		jErr := json.NewDecoder(r.Body).Decode(&allocation)
		if jErr != nil {
			http.Error(w, jErr.Error(), http.StatusBadRequest)
		}

		prorates, resp, aErr := util.AllocationCalculator(allocation.AllocationAmount, allocation.InvestorAmounts)
		if aErr != nil {
			println("Error calculating allocation: %v", aErr)
			http.Error(w, aErr.Error(), http.StatusInternalServerError)
			return
		}

		c.JSON(resp, prorates)

		server.Handle(w, r)
	})
}

//func main() {
//	app := fiber.New()
//
//	app.Use(cors.New(cors.Config{
//		AllowOrigins: "http://localhost:3000",
//		AllowHeaders: "Origin, Content-Type, Accept",
//		AllowMethods: "POST",
//	}))
//
//	opportunity := util.Allocate{}
//
//	app.Post("/api/invest", func(ctx *fiber.Ctx) error {
//		request := new(util.Allocate)
//
//		if pErr := ctx.BodyParser(request); pErr != nil {
//			return fmt.Errorf("failed to parse request body: %w", pErr)
//		}
//
//		opportunity = *request
//
//		prorates, _, pErr := util.AllocationCalculator(opportunity.AllocationAmount, opportunity.InvestorAmounts)
//		if pErr != nil {
//			return fmt.Errorf("error calculating prorate: %w", pErr)
//		}
//
//		return ctx.JSON(prorates)
//	})
//
//	log.Fatal(app.Listen(":4000"))
//}
