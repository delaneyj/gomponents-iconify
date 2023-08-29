package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartAreasplineVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 21H2V3h2v12.54L9.5 6L16 9.78l4.24-7.33l1.73 1L22 21Z"/>`),
		g.Group(children),
	)
}