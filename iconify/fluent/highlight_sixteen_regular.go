package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 1a.5.5 0 0 1 .5.5v3a.5.5 0 0 0 .5.5h9.002a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 1 1 0v3a1.5 1.5 0 0 1-1.001 1.415V7a2 2 0 0 1-2 2H11l.003 1.74a1.5 1.5 0 0 1-.69 1.265l-4.54 2.916a.5.5 0 0 1-.77-.421V9H5a2 2 0 0 1-2-2V5.915A1.5 1.5 0 0 1 2 4.5v-3a.5.5 0 0 1 .5-.5Zm3.503 8v4.585l3.77-2.422a.5.5 0 0 0 .23-.421L10 9H6.003ZM4 7a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V6H4v1Z"/>`),
		g.Group(children),
	)
}