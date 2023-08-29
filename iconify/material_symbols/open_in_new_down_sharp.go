package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInNewDownSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v9h-2V5H5v14h7v2H3Zm11 0v-2h3.6L8.3 9.7l1.4-1.4l9.3 9.275V14h2v7h-7Z"/>`),
		g.Group(children),
	)
}