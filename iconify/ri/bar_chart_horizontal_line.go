package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartHorizontalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3v2H3V3h9Zm4 16v2H3v-2h13Zm6-8v2H3v-2h19Z"/>`),
		g.Group(children),
	)
}