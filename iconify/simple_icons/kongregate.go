package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kongregate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.358 22.95v-3.736h1.551l.106-.141l-3.877-5.851l-3.172 3.264l-.026 2.351l.166.095l2.22 1.342l.071 2.676H.141l.053-3.021l2.027-.715l.106-.141V5.187l-.07-.141L0 4.165V.922h10.362v3.736h-2.22l-.106.141v7.014l7.472-6.802V4.87l-1.163-.352l-1.163-.352V.922h10.75v3.736h-1.41l-.352.106l-7.219 6.165l6.493 8.46l.246.246l2.31.787v2.656l-10.642-.128z"/>`),
		g.Group(children),
	)
}