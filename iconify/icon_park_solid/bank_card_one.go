package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBankCardOne0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M14 13V9a2 2 0 0 1 2-2h26a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2h-2"/><rect width="30" height="22" x="4" y="19" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M4 28h30"/><path stroke="#fff" d="M34 23v12M4 23v12"/><path stroke="#000" d="M11 34h8m6 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBankCardOne0)"/>`),
		g.Group(children),
	)
}