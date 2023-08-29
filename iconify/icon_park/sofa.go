package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sofa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M12 21H4V35H12V21Z"/><path fill="#2F88FF" d="M44 21H36V35H44V21Z"/><path stroke-linecap="round" d="M36 27H12V35H36V27Z"/><path stroke-linecap="round" d="M8 20V8H40V20"/><path stroke-linecap="round" d="M8 36V40"/><path stroke-linecap="round" d="M40 36V40"/></g>`),
		g.Group(children),
	)
}