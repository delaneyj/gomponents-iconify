package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3h6l3 7l-6 2l-6-2zm3 9L9 3m6 8l-3-8m0 16.5L9 21l.5-3.5l-2-2l3-.5l1.5-3l1.5 3l3 .5l-2 2L15 21z"/>`),
		g.Group(children),
	)
}