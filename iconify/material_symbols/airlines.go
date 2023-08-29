package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airlines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20L13 4h9l-3 16H2Zm12.5-6q1.05 0 1.775-.725T17 11.5q0-1.05-.725-1.775T14.5 9q-1.05 0-1.775.725T12 11.5q0 1.05.725 1.775T14.5 14Z"/>`),
		g.Group(children),
	)
}