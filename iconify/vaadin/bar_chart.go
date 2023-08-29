package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 15h15v1H0v-1zm0-4h3v3H0v-3zm4-2h3v5H4V9zm4-4h3v9H8V5zm4-5h3v14h-3V0z"/>`),
		g.Group(children),
	)
}