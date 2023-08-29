package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModelBuilderReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4z"/><path fill="currentColor" d="M22 4v3H10V4H2v8h8V9h7.023A4.946 4.946 0 0 0 16 12v16h2V12a2.998 2.998 0 0 1 2.997-3H22v3h8V4ZM8 10H4V6h4Zm20 0h-4V6h4Z"/>`),
		g.Group(children),
	)
}