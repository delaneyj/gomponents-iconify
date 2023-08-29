package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsDownCircleOneArrowKeyboardCircleButtonDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4v6M5.5 8.5L7 10l1.5-1.5"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}