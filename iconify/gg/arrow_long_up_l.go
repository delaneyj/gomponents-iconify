package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongUpL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.032 1.019l4.21 4.274l-1.424 1.404l-1.799-1.826l-.051 16.11h1.996v2h-6v-2h2.004l.051-16.157l-1.858 1.83l-1.403-1.425l4.274-4.21Z"/>`),
		g.Group(children),
	)
}