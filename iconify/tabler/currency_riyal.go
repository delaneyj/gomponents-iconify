package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRiyal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9v2a2 2 0 1 1-4 0v-1v1a2 2 0 1 1-4 0v-1v4a2 2 0 1 1-4 0v-2m15 .01V12m4-2v1a5 5 0 0 1-5 5"/>`),
		g.Group(children),
	)
}