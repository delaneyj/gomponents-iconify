package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="13" r="9" stroke-linejoin="round"/><path d="M4.5 44c0-6 7-16 19.5-16c0 0 2.759 0 5.782 1.09C32.744 30.16 36.5 31.149 36.5 28V7.75a3.75 3.75 0 1 1 7.5 0V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 44h44"/></g>`),
		g.Group(children),
	)
}