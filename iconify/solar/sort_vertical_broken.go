package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortVerticalBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 18v-2m0-10l4 4.125M16 6l-4 4.125M16 6v7m-8 5l4-4.125M8 18l-4-4.125M8 18v-7m0-5v2"/>`),
		g.Group(children),
	)
}