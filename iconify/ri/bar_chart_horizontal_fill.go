package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3v4H3V3h9Zm4 14v4H3v-4h13Zm6-7v4H3v-4h19Z"/>`),
		g.Group(children),
	)
}