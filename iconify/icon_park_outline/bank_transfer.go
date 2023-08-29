package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankTransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="10" rx="2"/><path stroke-linecap="round" d="M4 20h40M4 17v6m40-6v6m-15 6h8"/></g>`),
		g.Group(children),
	)
}