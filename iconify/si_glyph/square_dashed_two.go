package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareDashedTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 .834h5.083v5.041h.834V0H10v.834zM.834 5.875V.834h5.083V0H0v5.875h.834zm14.249 4.25v5.041H10V16h5.917v-5.875h-.834zm-9.166 5.041H.834v-5.041H0V16h5.917v-.834z"/>`),
		g.Group(children),
	)
}