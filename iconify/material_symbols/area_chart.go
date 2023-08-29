package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 16l-9.4-7.35l-3.975 5.475L3 10.5V7l4 3l5-7l5 4h4v9ZM3 20v-7l5 4l4-5.5l9 7.025V20H3Z"/>`),
		g.Group(children),
	)
}