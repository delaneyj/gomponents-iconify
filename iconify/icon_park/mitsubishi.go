package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mitsubishi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M17 19L24 30L31 19L24 8L17 19Z"/><path d="M30.6667 40L24 30H37.3333L44 40H30.6667Z"/><path d="M17.3333 40L24 30H10.6667L4 40H17.3333Z"/></g>`),
		g.Group(children),
	)
}