package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12h2v9H3v-9Zm16-4h2v13h-2V8Zm-8-6h2v19h-2V2Z"/>`),
		g.Group(children),
	)
}