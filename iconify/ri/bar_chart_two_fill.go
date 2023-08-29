package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13h6v8H2v-8ZM9 3h6v18H9V3Zm7 5h6v13h-6V8Z"/>`),
		g.Group(children),
	)
}