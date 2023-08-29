package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M32 22v2H0V0h2v22zM26 6l4 14H4v-9l7-9l9 9z"/>`),
		g.Group(children),
	)
}