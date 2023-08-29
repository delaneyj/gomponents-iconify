package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileQuality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M10 44H38C39.1046 44 40 43.1046 40 42V14H30V4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44Z"/><path stroke="#000" d="M30 4L40 14"/><path fill="#43CCF8" stroke="#fff" d="M19.2 21H28.8L32 25.1176L24 35L16 25.1176L19.2 21Z"/></g>`),
		g.Group(children),
	)
}