package hosts

import (
	"io"
	"log"
	"text/template"

	"github.com/Xuanwo/tiresias/config"
)

var t *template.Template

// Generate will generate the content.
func Generate(w io.Writer, s config.Server) (err error) {
	err = t.Execute(w, s)
	if err != nil {
		log.Fatalf("Template generate failed for %v", err)
	}
	return nil
}

func init() {
	var err error
	t, err = template.New("hosts").Parse(hostTemplate)
	if err != nil {
		log.Fatal(err)
	}
}
