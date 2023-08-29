package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareFourAnglePoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.959 3V.041H13v1.021H2.959V.041H0V3h1.011v10.041H0V16h2.959v-1.021H13V16h2.959v-2.959h-1.031V3h1.031zM13.938.938h1.084v1.125h-1.084V.938zm-13 0h1.084v1.125H.938V.938zm1.083 14.124H.937v-1.125h.073v.011h.917v-.011h.094v1.125zm13-1.124v1.125h-1.084v-1.125h.073v.011h.917v-.011h.094zm-1.01-.897H13v1.021H2.959v-1.021H1.928V3h1.031V1.979H13V3h1.011v10.041z"/>`),
		g.Group(children),
	)
}