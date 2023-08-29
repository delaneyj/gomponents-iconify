package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPenThreeContentCreationEditPenPensWrite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3a2.5 2.5 0 0 0-5 0v7.5l2.5 3l2.5-3Zm-5 1.5h5m2-2a2 2 0 0 1 4 0v8a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1Zm2 9v2"/><path d="M7.5 4.5h5a1 1 0 0 1 1 1v4"/></g>`),
		g.Group(children),
	)
}