package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FundsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.901 17.864l3.911-3.911l2.829 2.828l4.571-4.571l1.793 1.793v-5h-5l1.793 1.793l-3.157 3.157l-2.829-2.829l-4.945 4.946a9.965 9.965 0 0 1-.862-4.067c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10a9.987 9.987 0 0 1-8.104-4.14Z"/>`),
		g.Group(children),
	)
}