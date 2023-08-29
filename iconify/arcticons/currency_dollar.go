package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.52 30.877c1.08 1.404 2.432 1.928 4.314 1.928h2.605a4.394 4.394 0 0 0 4.389-4.399h0a4.394 4.394 0 0 0-4.39-4.399h-2.877a4.394 4.394 0 0 1-4.39-4.398h0a4.394 4.394 0 0 1 4.39-4.399h2.605c1.882 0 3.235.524 4.313 1.928M24 34.996V13.004"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}