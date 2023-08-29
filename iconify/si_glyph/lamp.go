package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.655.029H7.312c-.765 0-3.272 6.932-3.272 6.932h10.941S12.379.029 11.655.029zm-7.634 8v.929h4.981v5.084H6.396c-1.333 0-1.333 1.925-1.333 1.925h8.896s0-1.925-1.333-1.925H9.959V8.958h2.084v2h.875v-2h2.041v-.929H4.021z"/>`),
		g.Group(children),
	)
}