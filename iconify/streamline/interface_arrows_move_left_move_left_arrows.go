package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsMoveLeftMoveLeftArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 7h-6M3 4.5L.5 7L3 9.5"/><rect width="4.5" height="13" x="9" y=".5" rx="1"/></g>`),
		g.Group(children),
	)
}