package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HttpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 15V9h1.5v2h2V9H6v6H4.5v-2.5h-2V15H1Zm7.5 0v-4.5H7V9h4.5v1.5H10V15H8.5Zm5.5 0v-4.5h-1.5V9H17v1.5h-1.5V15H14Zm4 0V9h5v4h-3.5v2H18Zm1.5-3.5h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}