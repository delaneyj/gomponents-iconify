package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Segment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18v-2h12v2H9Zm0-5v-2h12v2H9ZM3 8V6h18v2H3Z"/>`),
		g.Group(children),
	)
}