package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tornado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.528 4A.514.514 0 0 0 5 4.5c0 .259.219.5.528.5H19v2H5.528C4.149 7 3 5.898 3 4.5S4.15 2 5.528 2H8v2H5.528ZM21 8v2H6V8h15ZM8 11h11v2H8v-2Zm-2 3h9v2H6v-2Zm-1 3h8v2H5v-2Zm3 3h7v2H8v-2Z"/>`),
		g.Group(children),
	)
}