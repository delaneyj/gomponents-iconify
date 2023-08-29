package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M69.8 325.8h139.6V512h93.1V325.8h139.6L256 93.1L69.8 325.8zM0 0v46.5h512V0H0z"/>`),
		g.Group(children),
	)
}