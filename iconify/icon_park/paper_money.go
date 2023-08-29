package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperMoney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M39 7H9C6.23858 7 4 9.23858 4 12V36C4 38.7614 6.23858 41 9 41H39C41.7614 41 44 38.7614 44 36V12C44 9.23858 41.7614 7 39 7Z"/><path stroke="#fff" stroke-linecap="round" d="M18 15L24 21L30 15"/><path stroke="#fff" stroke-linecap="round" d="M17 23H31"/><path stroke="#fff" stroke-linecap="round" d="M17 29H31"/><path stroke="#fff" stroke-linecap="round" d="M24 23V34"/></g>`),
		g.Group(children),
	)
}