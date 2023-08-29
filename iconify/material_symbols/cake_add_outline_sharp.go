package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeAddOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-8h2V8h6V6.55q-.45-.3-.725-.725T9 4.8q0-.375.15-.738T9.6 3.4L11 2l1.4 1.4q.3.3.45.662T13 4.8q0 .6-.275 1.025T12 6.55V8h6v6h2v8H2Zm4-8h10v-4H6v4Zm-2 6h14v-4H4v4Zm2-6h10H6Zm-2 6h14H4Zm14-6H4h14Zm1-6V6h-2V4h2V2h2v2h2v2h-2v2h-2Z"/>`),
		g.Group(children),
	)
}