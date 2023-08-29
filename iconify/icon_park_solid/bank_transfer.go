package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankTransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBankTransfer0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="10" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" stroke-linecap="round" d="M4 20h40"/><path stroke="#fff" stroke-linecap="round" d="M4 17v6m40-6v6"/><path stroke="#000" stroke-linecap="round" d="M29 29h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBankTransfer0)"/>`),
		g.Group(children),
	)
}