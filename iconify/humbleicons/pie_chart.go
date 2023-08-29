package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 15A9.004 9.004 0 0 1 12 21A9 9 0 0 1 9 3.512M12 3a9 9 0 0 1 9 9h-9V3Z"/>`),
		g.Group(children),
	)
}