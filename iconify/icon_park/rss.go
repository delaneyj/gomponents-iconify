package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M8 44L8 6C8 4.89543 8.89543 4 10 4H38C39.1046 4 40 4.89543 40 6V44L24 35.7273L8 44Z"/><path stroke="#fff" stroke-linecap="round" d="M23.9497 13.9497L23.9497 25.9497"/><path stroke="#fff" stroke-linecap="round" d="M17.9497 19.9497L29.9497 19.9497"/></g>`),
		g.Group(children),
	)
}