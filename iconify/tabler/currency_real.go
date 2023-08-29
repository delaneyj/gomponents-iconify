package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyReal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6h-4a3 3 0 0 0 0 6h1a3 3 0 0 1 0 6h-4M4 18V6h3a3 3 0 1 1 0 6H4c5.5 0 5 4 6 6m8-12V4m-1 16v-2"/>`),
		g.Group(children),
	)
}