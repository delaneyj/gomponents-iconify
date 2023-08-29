package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M13 24L24 14L35 24V44H13V24Z"/><path stroke="#000" stroke-linejoin="round" d="M7 8L10 4L13 8V44H7V8Z"/><path stroke="#000" stroke-linejoin="round" d="M35 8L38 4L41 8V44H35V8Z"/><path stroke="#fff" d="M24 25V35"/><path stroke="#fff" d="M20 29L28 29"/></g>`),
		g.Group(children),
	)
}