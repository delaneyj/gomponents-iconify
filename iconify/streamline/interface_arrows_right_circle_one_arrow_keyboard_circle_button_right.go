package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsRightCircleOneArrowKeyboardCircleButtonRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 7h6M8.5 5.5L10 7L8.5 8.5"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}