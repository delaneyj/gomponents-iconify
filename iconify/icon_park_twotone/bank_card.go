package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBankCard0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 10a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V10Z"/><path stroke-linecap="square" d="M4 16h40"/><path stroke-linecap="round" d="M27 32h9m8-22v16M4 10v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBankCard0)"/>`),
		g.Group(children),
	)
}