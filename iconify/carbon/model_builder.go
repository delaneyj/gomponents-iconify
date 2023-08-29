package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModelBuilder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 4v3H10V4H2v8h8V9h7.023A4.946 4.946 0 0 0 16 12v8a3.004 3.004 0 0 1-3 3h-3v-3H2v8h8v-3h3a5.006 5.006 0 0 0 5-5v-8a2.998 2.998 0 0 1 2.997-3H22v3h8V4ZM8 10H4V6h4Zm0 16H4v-4h4Zm20-16h-4V6h4Z"/>`),
		g.Group(children),
	)
}