package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBankCardTwo0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="M44 18V8H4v10"/><path fill="#fff" stroke="#fff" d="M44 18H4v22h40V18Z"/><path stroke="#000" stroke-linecap="round" d="M12 29h2m6 0h2m6 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBankCardTwo0)"/>`),
		g.Group(children),
	)
}