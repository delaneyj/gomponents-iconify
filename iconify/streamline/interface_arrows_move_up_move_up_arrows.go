package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsMoveUpMoveUpArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 6.5v-6M9.5 3L7 .5L4.5 3"/><rect width="4.5" height="13" x="4.75" y="4.75" rx="1" transform="rotate(90 7 11.25)"/></g>`),
		g.Group(children),
	)
}