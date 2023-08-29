package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18.5a6 6 0 0 1-5 0a6 6 0 0 0-5 .5a3 3 0 0 0 2-2.5V9a4 4 0 0 1 7.45-2m-2.55 6h-7"/>`),
		g.Group(children),
	)
}