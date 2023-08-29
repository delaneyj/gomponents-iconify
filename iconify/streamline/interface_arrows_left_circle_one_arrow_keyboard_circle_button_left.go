package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsLeftCircleOneArrowKeyboardCircleButtonLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 7H4m1.5-1.5L4 7l1.5 1.5"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}