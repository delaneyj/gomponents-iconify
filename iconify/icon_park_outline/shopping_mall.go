package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingMall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M8 44V6a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v38" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 44V6a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v38m0-29l10 6v23"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M4 44h40"/></g>`),
		g.Group(children),
	)
}