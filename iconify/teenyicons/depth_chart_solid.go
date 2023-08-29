package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DepthChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1 0H0v15h15V0h-1v5h-2v2h-1v3H9v2H8v-1H7V8H5V5H3V4H1V0Z"/>`),
		g.Group(children),
	)
}