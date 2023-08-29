package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsLeftCircleTwoArrowKeyboardCircleButtonLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8 4L5 7l3 3"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}