package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfloatLandscapeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v7h-7v9ZM6 8v6h2v-2.575l3.075 3.075l1.425-1.425L9.4 10H12V8Zm11 12v-7h5v7Z"/>`),
		g.Group(children),
	)
}