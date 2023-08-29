package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildrenPyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15 17h18v8H15zm-5 8h28v8H10v-8Zm-5 8h38v8H5v-8Zm19-16V7m5 0H19"/>`),
		g.Group(children),
	)
}