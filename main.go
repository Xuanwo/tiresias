package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"sort"

	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/constants"
	"github.com/Xuanwo/tiresias/destination"
	"github.com/Xuanwo/tiresias/destination/hosts"
	"github.com/Xuanwo/tiresias/destination/ssh"
	"github.com/Xuanwo/tiresias/model"
	"github.com/Xuanwo/tiresias/source"
	"github.com/Xuanwo/tiresias/source/fs"
)

func main() {
	app := cli.NewApp()
	app.Name = constants.Name
	app.Usage = constants.Usage
	app.Version = constants.Version

	// Setup flags.
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(c *cli.Context) error {
		conf, _ := config.New()
		err := conf.LoadFromFilePath(c.String("config"))
		if err != nil {
			return err
		}

		s := []model.Server{}
		sources := make([]source.Source, len(conf.Src))
		destinations := make([]destination.Destnation, len(conf.Dst))

		for k, v := range conf.Src {
			var src source.Source
			switch v.Type {
			case constants.TypeFs:
				src = &fs.Fs{}
				err = src.Init(v)
				if err != nil {
					return err
				}
			}
			sources[k] = src
		}

		for k, v := range conf.Dst {
			var dst destination.Destnation
			switch v.Type {
			case constants.TypeHosts:
				dst = &hosts.Hosts{}
				err = dst.Init(v)
				if err != nil {
					return err
				}
			case constants.TypeSSHConfig:
				dst = &ssh.SSH{}
				err = dst.Init(v)
				if err != nil {
					return err
				}
			}
			destinations[k] = dst
		}

		for _, v := range sources {
			ts, err := v.List()
			if err != nil {
				return err
			}
			s = append(ts)
		}

		for _, v := range destinations {
			_, err := v.Write(s...)
			if err != nil {
				return err
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
