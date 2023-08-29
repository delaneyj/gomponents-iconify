package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FundsBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.005 3.003h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1Zm11.793 6.793l-2.45 2.45l-2.121-2.122l-4.243 4.243l1.414 1.414l2.829-2.828l2.121 2.121l3.864-3.864l1.793 1.793v-5h-5l1.793 1.793Z"/>`),
		g.Group(children),
	)
}