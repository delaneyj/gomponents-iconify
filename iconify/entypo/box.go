package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.399 2H1.6c-.332 0-.6.267-.6.6V5h18V2.6a.6.6 0 0 0-.601-.6zM2 16.6c0 .77.629 1.4 1.399 1.4h13.2c.77 0 1.4-.631 1.4-1.4V6H2v10.6zM7 8h6v2H7V8z"/>`),
		g.Group(children),
	)
}