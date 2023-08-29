package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Verizon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.401 0h4.932v.005L14.234 32h-4.021L2.666 16h4.969l4.599 9.781z"/>`),
		g.Group(children),
	)
}