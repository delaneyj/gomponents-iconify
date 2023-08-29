package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#2A5ADA"/><path fill="#FFF" d="m16 6l-1.799 1.055L9.3 9.945L7.5 11v10l1.799 1.055l4.947 2.89L16.045 26l1.799-1.055l4.857-2.89L24.5 21V11l-1.799-1.055l-4.902-2.89L16 6zm-4.902 12.89v-5.78L16 10.22l4.902 2.89v5.78L16 21.78l-4.902-2.89z"/></g>`),
		g.Group(children),
	)
}