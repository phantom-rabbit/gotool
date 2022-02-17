package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("stoss-tools")

func main() {
	level := os.Getenv("LOTUS_LOG_LEVEL")
	if level != "" {
		logging.SetLogLevel("*", level)
	} else {
		logging.SetLogLevel("*", "INFO")
	}

	app := &cli.App{
		Name:    "stoss-tools",
		Usage:   "lotus-ctrl",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus",
			},
			&cli.StringFlag{
				Name:    "miner-repo",
				EnvVars: []string{"LOTUS_MINER_PATH"},
				Hidden:  true,
				Value:   "~/.lotusminer",
			},
		},
		Commands: []*cli.Command{

		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}

package main

import (
	"github.com/urfave/cli/v2"
)

var erasureSetCmd = &cli.Command{
	Name:  "time-epoch",
	Usage: "Make time transform epoch",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "time",
			Usage: "time format e.g 2006-01-02 15:04:05",
		},
	},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

