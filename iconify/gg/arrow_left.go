package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.328 11v2H7.5l3.243 3.243l-1.415 1.414L3.672 12l5.656-5.657l1.415 1.414L7.5 11h12.828Z"/>`),
		g.Group(children),
	)
}