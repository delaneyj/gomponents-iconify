package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.634 26L17.79 5.105a2 2 0 0 0-3.588.021L4.366 26H2v2h28v-2ZM15.99 5.979L20.947 16.5L19 17.798l-3-2l-3 2l-1.955-1.303Zm-5.805 12.346L13 20.202l3-2l3 2l2.81-1.873L25.422 26H6.575Z"/>`),
		g.Group(children),
	)
}