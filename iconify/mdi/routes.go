package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Routes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 10H5L3 8l2-2h6V3l1-1l1 1v1h6l2 2l-2 2h-6v2h6l2 2l-2 2h-6v6a2 2 0 0 1 2 2H9a2 2 0 0 1 2-2V10Z"/>`),
		g.Group(children),
	)
}