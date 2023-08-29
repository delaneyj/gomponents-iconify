package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.53 24a4.399 4.399 0 0 1 0 8.797h-7.258V15.203h7.257a4.399 4.399 0 0 1 0 8.797h0Zm-.001 0h-7.257m0 8.797h-2.2m2.2-17.594h-2.2m4.399 19.793v-2.199m4.398 2.199v-2.199m-4.398-17.594v-2.199m4.398 2.199v-2.199"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}