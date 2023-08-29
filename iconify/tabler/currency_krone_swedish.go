package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKroneSwedish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6v12m0-6c3.5 0 6-3 6-6m-6 6c3.5 0 6 3 6 6m4-8v8m4-8a4 4 0 0 0-4 4"/>`),
		g.Group(children),
	)
}