package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsUpCircleOneArrowUpKeyboardCircleButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 10V4M5.5 5.5L7 4l1.5 1.5"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}