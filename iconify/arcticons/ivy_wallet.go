package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IvyWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.312 38.862H11.688A6.318 6.318 0 0 1 5.5 33.817h0a48.707 48.707 0 0 1 0-19.634h0a6.318 6.318 0 0 1 6.188-5.045h24.624a6.318 6.318 0 0 1 6.188 5.045h0a48.707 48.707 0 0 1 0 19.634h0a6.318 6.318 0 0 1-6.188 5.045Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 18L24 30l7.5-12"/>`),
		g.Group(children),
	)
}