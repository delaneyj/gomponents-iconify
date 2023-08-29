package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18.89h1.414l9.314-9.314l-1.414-1.414L5 17.476v1.414Zm16 2H3v-4.243L16.435 3.212a1 1 0 0 1 1.414 0l2.829 2.829a1 1 0 0 1 0 1.414L9.243 18.89H21v2ZM15.728 6.748l1.414 1.414l1.414-1.414l-1.414-1.414l-1.414 1.414Z"/>`),
		g.Group(children),
	)
}