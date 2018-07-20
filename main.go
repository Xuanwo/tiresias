package main

import (
	"log"
	"os"
	"sort"

	"gopkg.in/urfave/cli.v1"

	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/constants"
	"github.com/Xuanwo/tiresias/contexts"
	"github.com/Xuanwo/tiresias/destination"
	"github.com/Xuanwo/tiresias/model"
	"github.com/Xuanwo/tiresias/source"
)

var (
	// Conf stores the global config.
	Conf *config.Config

	// ExpectedSources holds all expected sources
	ExpectedSources map[string]struct{}

	// StoredSources holds all stored source.
	StoredSources []string

	// AvailableSources holds all available sources
	AvailableSources []source.Source

	// Destinations holds all destinations
	Destinations []destination.Destination
)

func run(c *cli.Context) error {
	if len(c.Args()) == 0 {
		cli.ShowAppHelpAndExit(c, 0)
	}

	Conf, _ = config.New()
	err := Conf.LoadFromFilePath(c.String("config"))
	if err != nil {
		return err
	}

	err = contexts.SetupContexts(Conf)
	if err != nil {
		log.Fatalf("Contexts setup failed for %v.", err)
	}
	defer contexts.DB.Close()

	err = setup()
	if err != nil {
		log.Fatalf("Setup failed for %v.")
	}

	// Save servers to db.
	for _, v := range AvailableSources {
		err = source.SaveServers(v)
		if err != nil {
			return err
		}
	}

	// Load servers from db.
	for _, v := range Destinations {
		err = destination.LoadServers(v)
		if err != nil {
			return err
		}
	}

	// Delete servers in db but not expected.
	for _, v := range StoredSources {
		if _, ok := ExpectedSources[v]; !ok {
			err = model.DeleteSource(v)
			if err != nil {
				log.Printf("Delete source failed for %v.", err)
				continue
			}
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = constants.Name
	app.Usage = constants.Usage
	app.Version = constants.Version
	app.Action = run

	// Setup flags.
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
