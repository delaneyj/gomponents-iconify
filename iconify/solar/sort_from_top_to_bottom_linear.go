package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortFromTopToBottomLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 16h9m-7-5h7M8 6h5"/><path stroke-linejoin="round" d="M17 4v16l3-4"/></g>`),
		g.Group(children),
	)
}