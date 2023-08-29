package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetailsLess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9a1 1 0 0 0 0 2h18a1 1 0 1 0 0-2H3Zm0 4a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2H3Z"/>`),
		g.Group(children),
	)
}