package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellsSplitTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 6h7v-.5A2.5 2.5 0 0 0 14.5 3H10v3Zm7 7V7H3v6h14Zm-7-5v4H9V8h1Zm0 9h4.5a2.5 2.5 0 0 0 2.5-2.5V14h-7v3Zm-1-3H3v.5A2.5 2.5 0 0 0 5.5 17H9v-3Zm0-8V3H5.5A2.5 2.5 0 0 0 3 5.5V6h6Z"/>`),
		g.Group(children),
	)
}