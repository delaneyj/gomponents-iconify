package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HlsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15V9h1.5v2h2V9H8v6H6.5v-2.5h-2V15H3Zm12.5 0v-2H17v.5h2v-1h-3.5V9h5v2H19v-.5h-2v1h3.5V15h-5ZM10 15V9h1.5v4.5H14V15h-4Z"/>`),
		g.Group(children),
	)
}