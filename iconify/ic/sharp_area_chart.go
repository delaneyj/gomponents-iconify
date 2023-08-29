package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAreaChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13v7h18v-1.5l-9-7L8 17l-5-4zm0-6l4 3l5-7l5 4h4v8.97l-9.4-7.31l-3.98 5.48L3 10.44V7z"/>`),
		g.Group(children),
	)
}