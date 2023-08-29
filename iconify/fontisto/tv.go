package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M27.75 0H2.25A2.25 2.25 0 0 0 0 2.25v15a2.25 2.25 0 0 0 2.25 2.25h11.489v1.499H6.24a1.5 1.5 0 0 0 0 3h18a1.5 1.5 0 0 0 0-3h-7.501v-1.5H27.75a2.25 2.25 0 0 0 2.25-2.25v-15a2.25 2.25 0 0 0-2.25-2.25zM27 16.5H3V3h24z"/>`),
		g.Group(children),
	)
}