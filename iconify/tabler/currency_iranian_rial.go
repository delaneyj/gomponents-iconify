package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyIranianRial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4v9a2 2 0 0 1-2 2H6a3 3 0 0 1-3-3v-1m9-6v8a1 1 0 0 0 1 1h1a2 2 0 0 0 2-2v-1m5 3v1.096a5 5 0 0 1-3.787 4.85L17 20m-6-2h.01M14 18h.01"/>`),
		g.Group(children),
	)
}