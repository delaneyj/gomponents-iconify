package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundScatterPlot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="7" cy="14" r="3" fill="currentColor"/><circle cx="11" cy="6" r="3" fill="currentColor"/><circle cx="16.6" cy="17.6" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}