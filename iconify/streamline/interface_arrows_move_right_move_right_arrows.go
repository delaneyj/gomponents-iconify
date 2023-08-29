package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsMoveRightMoveRightArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 7h6M11 4.5L13.5 7L11 9.5"/><rect width="4.5" height="13" x=".5" y=".5" rx="1" transform="rotate(-180 2.75 7)"/></g>`),
		g.Group(children),
	)
}