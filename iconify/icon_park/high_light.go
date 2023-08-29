package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M6 44L6 25H12V17H36V25H42V44H6Z"/><path stroke-linecap="round" d="M17 17V8L31 4V17"/></g>`),
		g.Group(children),
	)
}