package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GanttChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 0h1v14h14v1H0V0Zm2 2h3v1H2V2Zm1 3h5v1H3V5Zm2 3h3v1H5V8Zm3 3h7v1H8v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}