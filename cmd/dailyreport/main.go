package main

import (
	"log"
	"os"

	"github.com/ShibataTakao/daily-report/internal/dailyreport"
	"github.com/urfave/cli/v2"
)

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:  "dailyreport",
		Usage: "Dailyreport manager",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create today's daily report",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Path to daily report directory",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_PATH"},
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Create(c.String("path"))
				},
			},
			{
				Name:  "validate",
				Usage: "Validate worktime in today's daily report",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Path to daily report directory",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_PATH"},
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Validate(c.String("path"))
				},
			},
			{
				Name:  "aggregate",
				Usage: "Aggregate tasks in some daily reports",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Path to daily report directory",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_PATH"},
					},
					&cli.StringFlag{
						Name:     "from",
						Usage:    "Beginning of date range. Format must be yyyymmdd.",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "to",
						Usage:    "End of date range. Format must be yyyymmdd.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Aggregate(c.String("path"), c.String("from"), c.String("to"))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
