package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c8.837 0 16 7.163 16 16s-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0zm0 6l-1.799 1.055L9.3 9.945L7.5 11v10l1.799 1.055l4.947 2.89L16.045 26l1.799-1.055l4.857-2.89L24.5 21V11l-1.799-1.055l-4.902-2.89L16 6zm0 4.22l4.902 2.89v5.78L16 21.78l-4.902-2.89v-5.78L16 10.22z"/>`),
		g.Group(children),
	)
}