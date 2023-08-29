package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M209.5 418.9h93.1V232.7h139.6L256 0L69.8 232.7h139.6v186.2zM0 465.5V512h512v-46.5H0z"/>`),
		g.Group(children),
	)
}