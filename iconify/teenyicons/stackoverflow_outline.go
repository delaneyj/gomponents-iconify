package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackoverflowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.5 9v5.5h10V9M4 12.5h7M4 10l7 .7M4.3 7.5L11 9M5.3 4.5l6.1 3.1M7 2l5 4.5M9.5.5l3.5 5"/>`),
		g.Group(children),
	)
}