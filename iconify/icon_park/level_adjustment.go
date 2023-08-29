package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelAdjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M11.1139 18L14.877 3.95581L41.9229 11.2027L40.1016 18M7.89848 30L6.07715 36.7973L33.1231 44.0442L36.8862 30"/><path stroke-dasharray="2 6" d="M4 24H44"/></g>`),
		g.Group(children),
	)
}