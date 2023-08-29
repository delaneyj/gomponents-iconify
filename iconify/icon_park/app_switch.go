package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M34 4H14V44H34V4Z"/><path d="M44 8H34V40H44V8Z"/><path d="M14 8H4V40H14V8Z"/></g>`),
		g.Group(children),
	)
}