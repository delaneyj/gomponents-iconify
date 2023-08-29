package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TetragonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.61 2.576a1.966 1.966 0 0 1 2.78 0l8.034 8.034a1.966 1.966 0 0 1 0 2.78l-8.034 8.034a1.966 1.966 0 0 1-2.78 0L2.576 13.39a1.966 1.966 0 0 1 0-2.78l8.034-8.034Z"/>`),
		g.Group(children),
	)
}