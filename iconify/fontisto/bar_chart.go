package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 12v8H6v-8zm6-8v16h-4V4zm16 18v2H0V0h2v22zM22 8v12h-4V8zm6-6v18h-4V2z"/>`),
		g.Group(children),
	)
}