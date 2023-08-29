package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v8h15V2H2zm17 1a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v1zm-5.212 3.358a1 1 0 0 1-1.281.597l-3.76-1.368a1 1 0 0 1 .685-1.88l3.759 1.369a1 1 0 0 1 .597 1.282zM10.03 7.989a1 1 0 0 1-1.282.598L4.989 7.22a1 1 0 0 1 .684-1.88l3.759 1.369a1 1 0 0 1 .598 1.281zm.06-3.342v3a1 1 0 1 1-2 0v-3a1 1 0 1 1 2 0z"/>`),
		g.Group(children),
	)
}