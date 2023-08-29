package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposureSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h14V5L5 19Zm9.5-1v-2h-2v-1.5h2v-2H16v2h2V16h-2v2h-1.5ZM6 8.5h5V7H6v1.5Z"/>`),
		g.Group(children),
	)
}