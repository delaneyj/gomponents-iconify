package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M34 10L42 18"/><path d="M42 10L34 18"/><path d="M44 30L37 38L33 34"/><path fill="#2F88FF" d="M26 10H4V18H26V10Z"/><path fill="#2F88FF" d="M26 30H4V38H26V30Z"/></g>`),
		g.Group(children),
	)
}