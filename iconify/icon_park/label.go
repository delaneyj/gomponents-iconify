package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9 4H39V44L24 33.4286L9 44V4Z"/><rect width="30" height="12" x="9" y="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}