package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 17H15l3.5-5L15 7H6.5l3.5 5l-3.5 5m8.5 2H3l4.5-7L3 5h12c.69 0 1.23.3 1.64.86L21 12l-4.36 6.14c-.41.56-.95.86-1.64.86Z"/>`),
		g.Group(children),
	)
}