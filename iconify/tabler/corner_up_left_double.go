package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeftDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 18v-6a3 3 0 0 0-3-3H9"/><path d="M13 13L9 9l4-4m-5 8L4 9l4-4"/></g>`),
		g.Group(children),
	)
}