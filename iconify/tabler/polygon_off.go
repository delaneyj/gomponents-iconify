package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolygonOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0m7 3a2 2 0 1 0 4 0a2 2 0 1 0-4 0M3 11a2 2 0 1 0 4 0a2 2 0 1 0-4 0m10 8a2 2 0 1 0 4 0a2 2 0 1 0-4 0M6.5 9.5l1.546-1.311M14 5.5L17 7m1.5 3l-1.185 3.318m-1.062 2.972L16 17m-2.5.5l-7-5M3 3l18 18"/>`),
		g.Group(children),
	)
}