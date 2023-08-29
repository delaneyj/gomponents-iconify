package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2 4.706C2 14.257 9.743 22 19.294 22L23 18.294l-4.323-4.323l-3.089 3.088l-8.647-8.647l3.088-3.088L5.706 1L2 4.706Z"/>`),
		g.Group(children),
	)
}