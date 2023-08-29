package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsRightCircleTwoArrowKeyboardCircleButtonRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6 4l3 3l-3 3"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}