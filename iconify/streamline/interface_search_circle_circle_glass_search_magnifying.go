package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSearchCircleCircleGlassSearchMagnifying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><circle cx="6.5" cy="6.5" r="2.5"/><path d="M10 10L8.27 8.27"/></g>`),
		g.Group(children),
	)
}