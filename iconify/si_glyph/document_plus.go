package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.927.062L4.073.114l-.021 9.833l1.969.015v2.017H8v3.083H6.062V16h9.907V6.012H9.927V.062z"/><path d="M11 0v4.969h5L11 0zM6.979 13H4.992v-2h-.963v2H2.021v.986h2.008v2.008h.963v-2.008h1.987V13z"/></g>`),
		g.Group(children),
	)
}