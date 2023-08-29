package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsDownCircleTwoArrowKeyboardCircleButtonDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 6L7 9L4 6"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}