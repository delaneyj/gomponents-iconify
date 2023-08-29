package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noodles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" d="M4 24C4 35.0457 10.6667 44 24 44C37.3333 44 44 35.0457 44 24H4Z" clip-rule="evenodd"/><path d="M16 24V8"/><path d="M24 24V6"/><path d="M32 24V4"/><path d="M8 24V10"/><path d="M4 13L44 4"/></g>`),
		g.Group(children),
	)
}