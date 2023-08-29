package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineTwoAnglePoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.938.938h1.084v1.125H7.938V.938zm1.083 14.124H7.937v-1.125h.073v.011h.917v-.011h.094v1.125zm.938-2.021H8.928V3h1.031V.041H7V3h1.011v10.041H7V16h2.959v-2.959z"/>`),
		g.Group(children),
	)
}