package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundBubbleChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="7.2" cy="14.4" r="3.2" fill="currentColor"/><circle cx="14.8" cy="18" r="2" fill="currentColor"/><circle cx="15.2" cy="8.8" r="4.8" fill="currentColor"/>`),
		g.Group(children),
	)
}