package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 7.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0z"/>`),
		g.Group(children),
	)
}