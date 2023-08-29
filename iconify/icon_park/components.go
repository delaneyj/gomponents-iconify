package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Components(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 11L24 4L31 11L24 18L17 11Z"/><path d="M30 25L37 18L44 25L37 32L30 25Z"/><path d="M17 37L24 30L31 37L24 44L17 37Z"/><path d="M4 24L11 17L18 24L11 31L4 24Z"/></g>`),
		g.Group(children),
	)
}