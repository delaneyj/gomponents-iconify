package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsRoundRightDiagramRoundArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m9.25.5l2.5 2.5l-2.5 2.5"/><path d="M12.25 8.25A5.25 5.25 0 1 1 7 3h4.75"/></g>`),
		g.Group(children),
	)
}