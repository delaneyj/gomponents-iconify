package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyShekel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 18V6h4a4 4 0 0 1 4 4v4"/><path d="M18 6v12h-4a4 4 0 0 1-4-4v-4"/></g>`),
		g.Group(children),
	)
}