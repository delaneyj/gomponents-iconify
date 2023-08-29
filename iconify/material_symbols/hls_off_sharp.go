package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HlsOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.825 15l-2-2H17v.5h2v-1h-3.5V9h5v2H19v-.5h-2v1h3.5V15h-2.675ZM3 15V9h1.5v2h2V9H8v6H6.5v-2.5h-2V15H3Zm16.775 7.6L1.4 4.225L2.8 2.8l18.375 18.4l-1.4 1.4ZM10 15v-3.575l1.5 1.5v.575h.6l1.5 1.5H10Z"/>`),
		g.Group(children),
	)
}