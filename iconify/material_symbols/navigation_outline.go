package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 21l-1-1l8-18l8 18l-1 1l-7-3l-7 3Zm2.1-3.1l4.9-2.1l4.9 2.1l-4.9-11l-4.9 11Zm4.9-2.1Z"/>`),
		g.Group(children),
	)
}