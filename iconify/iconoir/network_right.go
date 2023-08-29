package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 22 21)"/><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 7 15.5)"/><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 22 10)"/><path d="M17 17.5h-3.5a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2H17M11.5 12H7"/></g>`),
		g.Group(children),
	)
}