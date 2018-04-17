package main

import (
	"github.com/go-spatial/go-wfs/wfs3"
	"fmt"
)

func main() {
	r := &wfs3.RootContent{
		Links: []*wfs3.Link{
			&wfs3.Link{Href: "http://somesite.com/path/to/stuff"},
			&wfs3.Link{Href: "http://somesite.com/path/to/otherstuff"},
		},
	}
	rslt := wfs3.DemoRootContentPageCreation(r, "A title")
	fmt.Println(rslt)
}
