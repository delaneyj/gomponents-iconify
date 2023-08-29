package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 .016a7.5 7.5 0 1 0 5.438 13.13L7.15 7.857A.5.5 0 0 1 7 7.5V.016Z"/><path fill="currentColor" d="M13.145 12.438A7.47 7.47 0 0 0 15 7.5a7.476 7.476 0 0 0-.581-2.9L8.344 7.637l4.801 4.801Zm.825-8.732A7.499 7.499 0 0 0 8 .016v6.675l5.97-2.985Z"/>`),
		g.Group(children),
	)
}