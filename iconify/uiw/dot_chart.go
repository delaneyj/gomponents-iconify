package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.02 13.997a1 1 0 0 1 1 1V17a1 1 0 0 1-1 1.001H1a1 1 0 0 1-1-1v-2.002a1 1 0 0 1 1-1.001h13.02ZM19 8.007a1 1 0 0 1 1 1.002v2.001a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1.001h18ZM11.03 2a1 1 0 0 1 1 1v2.002a1 1 0 0 1-1 1.001H1.001a1 1 0 0 1-1-1V3A1 1 0 0 1 1 2h10.03Z"/>`),
		g.Group(children),
	)
}