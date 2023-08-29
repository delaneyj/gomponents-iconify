package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOneBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-3.525-3.55q.675-.675 1.575-1.063T12 16q1.05 0 1.95.388t1.575 1.062L12 21Z"/>`),
		g.Group(children),
	)
}