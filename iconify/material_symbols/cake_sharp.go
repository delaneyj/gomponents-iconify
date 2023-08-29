package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-7h18v7H3Zm2-9V8h6V6.55q-.45-.3-.725-.725T10 4.8q0-.375.15-.738t.45-.662L12 2l1.4 1.4q.3.3.45.662T14 4.8q0 .6-.275 1.025T13 6.55V8h6v5H5Z"/>`),
		g.Group(children),
	)
}