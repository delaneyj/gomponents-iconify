package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEuro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.525 30.695a5.888 5.888 0 0 1-4.516 2.102h0a5.9 5.9 0 0 1-5.9-5.9v-5.794a5.9 5.9 0 0 1 5.9-5.9h0a5.89 5.89 0 0 1 4.535 2.125M15.91 21.862h6.598M15.91 26.18h6.598"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}