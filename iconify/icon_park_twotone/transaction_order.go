package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransactionOrder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTransactionOrder0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="30" height="36" x="9" y="8" fill="#555" rx="2"/><path stroke-linecap="round" d="M18 4v6m12-6v6m-14 9h16m-16 8h12m-12 8h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTransactionOrder0)"/>`),
		g.Group(children),
	)
}