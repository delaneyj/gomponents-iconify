package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.29 12.296a1 1 0 1 0 1.423 1.406l7.289-7.376v17.675a1 1 0 1 0 2 0V6.328l7.286 7.374a1 1 0 0 0 1.423-1.405L14.89 3.368a1.25 1.25 0 0 0-1.778 0L4.29 12.297Z"/>`),
		g.Group(children),
	)
}