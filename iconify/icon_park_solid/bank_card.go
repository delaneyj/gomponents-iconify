package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBankCard0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 10a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V10Z"/><path stroke="#000" stroke-linecap="square" d="M4 16h40"/><path stroke="#000" stroke-linecap="round" d="M27 32h9"/><path stroke="#fff" stroke-linecap="round" d="M44 10v16M4 10v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBankCard0)"/>`),
		g.Group(children),
	)
}