package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 9a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H5Zm0 4a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H5Z"/>`),
		g.Group(children),
	)
}