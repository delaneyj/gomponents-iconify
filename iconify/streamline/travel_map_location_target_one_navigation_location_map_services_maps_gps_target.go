package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapLocationTargetOneNavigationLocationMapServicesMapsGpsTarget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="4.5"/><circle cx="7" cy="7" r=".5"/><path d="M7 2.5v-2m0 13v-2M11.5 7h2M.5 7h2"/></g>`),
		g.Group(children),
	)
}