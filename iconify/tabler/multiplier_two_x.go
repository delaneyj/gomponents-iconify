package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiplierTwoX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 16l4-4m0 4l-4-4m-8-2a2 2 0 1 1 4 0c0 .591-.417 1.318-.816 1.858L6 16.001h4"/>`),
		g.Group(children),
	)
}