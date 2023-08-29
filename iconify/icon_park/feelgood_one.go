package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeelgoodOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 6H4V36H25L35 41V36H44V6Z"/><path stroke="#fff" d="M13 23C13 23 17 27 24 27C31 27 35 23 35 23"/></g>`),
		g.Group(children),
	)
}