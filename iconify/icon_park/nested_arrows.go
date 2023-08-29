package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestedArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 17V4H4V30H17"/><path fill="#2F88FF" d="M43 43V17H17V43H43Z"/><path d="M33 30H17M33 30L28 25L33 30ZM33 30L28 35L33 30Z"/><path d="M17 17V43"/></g>`),
		g.Group(children),
	)
}