package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 9.9l3 2.1l-3 2.1V9.9m-9 0L9 12l-3 2.1V9.9M13 6v12l8.5-6L13 6M4 6v12l8.5-6L4 6Z"/>`),
		g.Group(children),
	)
}