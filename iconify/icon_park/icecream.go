package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icecream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path d="M33 18H15L21 40C21 40 22 43 24 43C26 43 27 40 27 40L33 18Z"/><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M36 18H12C12 10 17 4 24 4C31 4 36 10 36 18Z"/></g>`),
		g.Group(children),
	)
}