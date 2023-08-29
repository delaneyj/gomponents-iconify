package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsMoveDownMoveDownArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 7.5v6M4.5 11L7 13.5L9.5 11"/><rect width="4.5" height="13" x="4.75" y="-3.75" rx="1" transform="rotate(-90 7 2.75)"/></g>`),
		g.Group(children),
	)
}