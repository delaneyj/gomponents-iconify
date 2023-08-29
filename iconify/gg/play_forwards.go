package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayForwards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.002 17h-3V7h3v10Zm-4-5L10 17V7l7.002 5ZM2 17l7.002-5L2 7v10Z"/>`),
		g.Group(children),
	)
}