package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 .613l11 3.666V22H2V3h9V.613ZM11 5H4v15h10.838L11 18.721V5.001Zm9 14.613V13.72l-7-2.333v5.891l7 2.334Zm0-8V5.72l-7-2.333v5.891l7 2.334Z"/>`),
		g.Group(children),
	)
}