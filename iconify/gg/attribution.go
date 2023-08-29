package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attribution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 8a2 2 0 0 0 1.732-1H14a2 2 0 1 1 0 4h-4a4 4 0 0 0 0 8h6.268A2 2 0 0 0 20 18a2 2 0 0 0-3.732-1H10a2 2 0 1 1 0-4h4a4 4 0 0 0 0-8H7.732A2 2 0 1 0 6 8Z"/>`),
		g.Group(children),
	)
}