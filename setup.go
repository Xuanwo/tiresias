package main

import (
	"log"

	"github.com/Xuanwo/tiresias/constants"
	"github.com/Xuanwo/tiresias/destination"
	"github.com/Xuanwo/tiresias/destination/hosts"
	"github.com/Xuanwo/tiresias/destination/ssh"
	"github.com/Xuanwo/tiresias/model"
	"github.com/Xuanwo/tiresias/source"
	"github.com/Xuanwo/tiresias/source/consul"
	"github.com/Xuanwo/tiresias/source/fs"
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
