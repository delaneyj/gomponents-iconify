package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDashboardLayoutOneCornersDashboardFrameLayoutMatOctagonSquareTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M.5 6.5L6 .5m7.5 6L8 .5m-7.5 7l5.5 6m7.5-6l-5.5 6"/></g>`),
		g.Group(children),
	)
}