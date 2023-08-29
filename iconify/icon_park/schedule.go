package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schedule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="40" height="30" x="4" y="10" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="2"/><path stroke="#000" d="M14 6V14"/><path stroke="#fff" d="M25 23L14 23"/><path stroke="#fff" d="M34 31L14 31"/><path stroke="#000" d="M34 6V14"/></g>`),
		g.Group(children),
	)
}