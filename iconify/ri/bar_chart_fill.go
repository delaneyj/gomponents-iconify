package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12h4v9H3v-9Zm14-4h4v13h-4V8Zm-7-6h4v19h-4V2Z"/>`),
		g.Group(children),
	)
}