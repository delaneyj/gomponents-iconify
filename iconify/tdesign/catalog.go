package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Catalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h18v22H3V1Zm2 2v18h14V3H5Zm3 4h8v2H8V7Zm0 4h8v2H8v-2Zm0 4h8v2H8v-2Z"/>`),
		g.Group(children),
	)
}