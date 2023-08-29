package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12h12c0-6.627-5.373-12-12-12z"/><path fill="currentColor" d="M10.5 13.5V.105C4.552.873.002 5.905 0 12c0 6.627 5.373 12 12 12c6.095-.003 11.127-4.552 11.889-10.44l.006-.06z"/>`),
		g.Group(children),
	)
}