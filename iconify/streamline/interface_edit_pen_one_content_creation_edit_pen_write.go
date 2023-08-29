package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPenOneContentCreationEditPenWrite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m.5 13.5l2-2m2.21.79a1 1 0 0 1-1.42 0l-1.58-1.58a1 1 0 0 1 0-1.42l8.17-8.17a2.12 2.12 0 0 1 3 3Z"/><path d="M11.5 5.5L7.21 1.21a1 1 0 0 0-1.42 0L2.5 4.5"/></g>`),
		g.Group(children),
	)
}