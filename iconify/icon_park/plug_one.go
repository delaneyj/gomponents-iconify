package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="24" height="24" x="12" y="12" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 36V41C24 42.6569 22.6569 44 21 44H8"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M20 12V4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 12V4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M22 24H26"/></g>`),
		g.Group(children),
	)
}