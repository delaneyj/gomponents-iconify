package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M10 44H38C39.1046 44 40 43.1046 40 42V14H30V4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M30 4L40 14"/><circle cx="22" cy="26" r="6" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M27 30L32 34"/></g>`),
		g.Group(children),
	)
}