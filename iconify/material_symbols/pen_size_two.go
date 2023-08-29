package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenSizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.3 18.7q-.275-.275-.275-.7t.275-.7l12-12q.275-.3.688-.3t.712.3q.275.275.275.7t-.275.7l-12 12q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}