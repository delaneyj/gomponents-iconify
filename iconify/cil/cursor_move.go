package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m130.412 323.98l-51.883-51.882H240v161.47l-51.882-51.881l-22.628 22.626l90.51 90.51l90.51-90.51l-22.628-22.626L272 433.568v-161.47h160.667l-51.883 51.882l22.628 22.627l90.509-90.509l-90.509-90.51l-22.628 22.627l51.883 51.883H272V79.432l51.882 51.881l22.628-22.626L256 18.177l-90.51 90.51l22.628 22.626L240 79.432v160.666H78.529l51.883-51.883l-22.628-22.627l-90.51 90.51l90.511 90.51l22.627-22.628z"/>`),
		g.Group(children),
	)
}