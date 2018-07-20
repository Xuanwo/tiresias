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
	"github.com/Xuanwo/tiresias/destination/hosts"
	"github.com/Xuanwo/tiresias/destination/ssh"
	"github.com/Xuanwo/tiresias/model"
	"github.com/Xuanwo/tiresias/source"
	"github.com/Xuanwo/tiresias/source/consul"
	"github.com/Xuanwo/tiresias/source/fs"
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

func setup() (err error) {
	StoredSources, err = model.ListSources()
	if err != nil {
		log.Fatalf("List stored sources failed for %v.", err)
		return
	}

	// Setup expected and available sources.
	ExpectedSources = make(map[string]struct{})

	for _, v := range Conf.Src {
		var src source.Source
		switch v.Type {
		case constants.TypeFs:
			src = &fs.Fs{}

			err = source.LoadConfig(src, v)
			if err != nil {
				log.Printf("Source %s %s load config failed for %v.",
					v.Type, v.Options, err)
				continue
			}

			ExpectedSources[src.Name()] = struct{}{}

			err = src.Init()
			if err != nil {
				log.Printf("Source %s init failed for %v.",
					src.Name(), err)
				continue
			}
		case constants.TypeConsul:
			src = &consul.Consul{}

			err = source.LoadConfig(src, v)
			if err != nil {
				log.Printf("Source %s %s load config failed for %v.",
					v.Type, v.Options, err)
				continue
			}

			ExpectedSources[src.Name()] = struct{}{}

			err = src.Init()
			if err != nil {
				log.Printf("Source %s init failed for %v.",
					src.Name(), err)
				continue
			}
		default:
			log.Printf("Type %s is not supported.", v.Type)
			continue
		}

		AvailableSources = append(AvailableSources, src)
	}

	// Setup destinations
	for _, v := range Conf.Dst {
		var dst destination.Destination
		switch v.Type {
		case constants.TypeHosts:
			dst = &hosts.Hosts{}

			err = destination.LoadConfig(dst, v)
			if err != nil {
				log.Printf("Destination %s %s load config failed for %v.",
					v.Type, v.Options, err)
				continue
			}

			err = dst.Init()
			if err != nil {
				log.Printf("Destnation %s %s init failed for %v.",
					v.Type, v.Options, err)
				continue
			}
		case constants.TypeSSHConfig:
			dst = &ssh.SSH{}

			err = destination.LoadConfig(dst, v)
			if err != nil {
				log.Printf("Destination %s %s load config failed for %v.",
					v.Type, v.Options, err)
				continue
			}

			err = dst.Init()
			if err != nil {
				log.Printf("Destnation %s %s init failed for %v.",
					v.Type, v.Options, err)
				continue
			}
		default:
			log.Printf("Type %s is not supported.", v.Type)
			continue
		}
		Destinations = append(Destinations, dst)
	}

	return
}

func run(c *cli.Context) error {
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

	for _, v := range AvailableSources {
		err = source.SaveServers(v)
		if err != nil {
			return err
		}
	}

	for _, v := range Destinations {
		err = destination.LoadServers(v)
		if err != nil {
			return err
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
