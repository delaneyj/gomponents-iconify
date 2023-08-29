package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zorin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5.333 25.26l2.661 4.604h16.011l2.661-4.604zM32 16l-2.683 4.651H12.286l17.031-9.301zM0 16l2.683-4.651h17.031L2.683 20.65zm5.333-9.26l2.661-4.604h16.011l2.661 4.604z"/>`),
		g.Group(children),
	)
}