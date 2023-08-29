package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faceu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V9C42 7.34315 40.6569 6 39 6Z"/><path fill="#43CCF8" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M27 12V16C29.6667 16.8333 36 19 36 26C36 33 30 36 24 36C18 36 12 33 12 26C12 18 21.6667 13.1667 27 12Z"/><ellipse cx="29" cy="26" fill="#fff" rx="3" ry="4"/><ellipse cx="19" cy="26" fill="#fff" rx="3" ry="4"/></g>`),
		g.Group(children),
	)
}