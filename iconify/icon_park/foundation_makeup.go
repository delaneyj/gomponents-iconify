package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoundationMakeup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M13 19C13 17.8954 13.8954 17 15 17H31C32.1046 17 33 17.8954 33 19V38C33 38 33 44 27 44C21 44 25 44 19 44C13 44 13 38 13 38V19Z"/><rect width="10" height="5" x="18" y="12"/><path d="M18 12C18 10 18.0001 12 18 9C17.9999 6 21 4 24 4C27 4 35 4 35 4C35 4 28.0001 6 28 10C27.9999 14 28 12 28 12"/></g>`),
		g.Group(children),
	)
}