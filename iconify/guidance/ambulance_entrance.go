package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmbulanceEntrance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 6.5v-3H.5V20h3m11-13.5h4a7.138 7.138 0 0 0 4.184 5.637l.816.363V17a3 3 0 0 1-3 3m-6-13.5V13h9m-3 7a2.5 2.5 0 0 1-5 0m5 0a2.5 2.5 0 0 0-5 0m0 0h-7m-5 0a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0M6 6.5V9H3.5v3H6v2.5h3V12h2.5V9H9V6.5H6Z"/>`),
		g.Group(children),
	)
}