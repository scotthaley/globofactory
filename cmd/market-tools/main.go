package main

import (
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/scotthaley/globofactory/internal/database"
	"github.com/scotthaley/globofactory/pkg/entity"
	market_listing "github.com/scotthaley/globofactory/pkg/market-listing"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	database.InitDB()

	app := &cli.App{
		Name: "market-tools",
		Usage: "Let's manipulate the market!",
		Commands: []*cli.Command{
			{
				Name: "search-entities",
				Aliases: []string{"se"},
				Usage: "search-entities ore",
				Action: func(c *cli.Context) error {
					s := c.Args().Get(0)
					found := entity.SearchEntityTypes(s)
					fmt.Printf("searching for: %v\n", s)
					for i := range found {
						fmt.Printf("%v - %v\n", found[i].Display, found[i].Code)
					}
					return nil
				},
			},
			{
				Name: "create-listing",
				Aliases: []string{"cl"},
				Usage: "create-listing -c gold-ore -t sell -p 1.50 -a 200 -e 2021-05-05T12:00:00",
				Flags: []cli.Flag {
					&cli.StringFlag{
						Name: "code",
						Aliases: []string{"c"},
						Usage: "entity code name for listing (\"gold-or\")",
						Required: true,
					},
					&cli.StringFlag{
						Name: "type",
						Aliases: []string{"t"},
						Usage: "listing type (\"buy\" or \"sell\")",
						DefaultText: "\"buy\"",
					},
					&cli.Float64Flag{
						Name: "price",
						Aliases: []string{"p"},
						Usage: "listing price (\"450.32\")",
						DefaultText: "1.00",
					},
					&cli.Int64Flag{
						Name: "amount",
						Aliases: []string{"a"},
						Usage: "amount of items to list (\"500\")",
						DefaultText: "100",
					},
					&cli.TimestampFlag{
						Name: "expiry",
						Aliases: []string{"e"},
						Usage: "DateTime listing will expire",
						Layout: "2006-01-02T15:04:05",
						DefaultText: "no expiration",
					},
				},
				Action: func(c *cli.Context) error {
					code := c.String("code")
					oType := market_listing.Buy
					if c.String("type") == "sell" {
						oType = market_listing.Sell
					}
					price := c.Float64("price")
					if price == 0 {
						price = 1.00
					}
					amount := c.Int64("amount")
					if amount == 0 {
						amount = 100
					}
					expiry := pgtype.Timestamp{}
					expiry.Set(c.Timestamp("expiry"))

					now := pgtype.Timestamp{}
					now.Set(time.Now())

					market_listing.Create(market_listing.MarketListing{
						Code:      code,
						OrderType: oType,
						Price: price,
						Amount: amount,
						ListingDate: now,
						ExpiryDate: expiry,
					})
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
