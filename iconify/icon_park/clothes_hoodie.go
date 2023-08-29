package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesHoodie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37 22V37M11 37V44H37V37M11 37H4V22C4 19 6 15.5 9 13C12 10.5 16 10 16 10L24 18M11 37V22M37 37H44V22C44 19 42 15.5 39 13C36 10.5 32 10 32 10L24 18M24 18V27"/><path fill="#2F88FF" d="M9 13C12 10.5 16 10 16 10L24 18L32 10C32 10 36 10.5 39 13L41 8.5L39 4H9L7 8.5L9 13Z"/></g>`),
		g.Group(children),
	)
}