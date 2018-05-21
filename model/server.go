package model

// Server stores all infomation that user needed to access a server.
type Server struct {
	Name         string `yaml:"name"`
	Address      string `yaml:"address"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	IdentityFile string `yaml:"identity_file"`
}
