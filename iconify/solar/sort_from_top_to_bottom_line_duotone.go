package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortFromTopToBottomLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 16h9"/><path d="M6 11h7" opacity=".7"/><path d="M8 6h5" opacity=".3"/><path stroke-linejoin="round" d="M17 4v16l3-4"/></g>`),
		g.Group(children),
	)
}