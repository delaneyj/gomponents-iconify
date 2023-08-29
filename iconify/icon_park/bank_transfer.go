package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankTransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="10" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M4 20H44"/><path stroke="#000" stroke-linecap="round" d="M4 17V23"/><path stroke="#000" stroke-linecap="round" d="M44 17V23"/><path stroke="#fff" stroke-linecap="round" d="M29 29L37 29"/></g>`),
		g.Group(children),
	)
}