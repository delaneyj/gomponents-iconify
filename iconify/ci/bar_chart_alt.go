package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 21H2V11a2 2 0 0 1 2-2h4V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3h4a2 2 0 0 1 2 2v12ZM16 9v10h4V9h-4Zm-6-5v15h4V4h-4Zm-6 7v8h4v-8H4Z"/>`),
		g.Group(children),
	)
}