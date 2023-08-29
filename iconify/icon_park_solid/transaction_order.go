package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransactionOrder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTransactionOrder0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="30" height="36" x="9" y="8" fill="#fff" stroke="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M18 4v6m12-6v6"/><path stroke="#000" stroke-linecap="round" d="M16 19h16m-16 8h12m-12 8h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTransactionOrder0)"/>`),
		g.Group(children),
	)
}