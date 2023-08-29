package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linejoin="round" d="M15 12L16.2 5H31.8L33 12"/><path stroke="#000" stroke-linecap="round" d="M6 12H42"/><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M37 12L35 43H13L11 12H37Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M19 35H29"/></g>`),
		g.Group(children),
	)
}