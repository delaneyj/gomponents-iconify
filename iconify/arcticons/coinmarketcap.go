package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coinmarketcap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.935 39.466A21.499 21.499 0 1 1 45.5 24h0c-.182 8.576-9.348 10.014-10.25 1.656l-1.342-8.85l-11.192 14.817l-1.342-15.96l-14.96 20.7"/>`),
		g.Group(children),
	)
}