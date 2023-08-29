package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankTransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBankTransfer0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="10" fill="#555" rx="2"/><path stroke-linecap="round" d="M4 20h40M4 17v6m40-6v6m-15 6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBankTransfer0)"/>`),
		g.Group(children),
	)
}