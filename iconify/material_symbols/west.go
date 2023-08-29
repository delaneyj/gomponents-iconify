package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func West(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 19l-7-7l7-7l1.4 1.4L5.825 11H22v2H5.825l4.6 4.6L9 19Z"/>`),
		g.Group(children),
	)
}