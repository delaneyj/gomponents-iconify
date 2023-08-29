package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingMall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShoppingMall0"><g fill="none"><path fill="#555" fill-rule="evenodd" d="M8 44V6a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v38" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 44V6a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v38m0-29l10 6v23"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M4 44h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShoppingMall0)"/>`),
		g.Group(children),
	)
}