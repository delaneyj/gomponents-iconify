package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" stroke-linejoin="round" d="M24 5L2 43H46L24 5Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M24 35V36"/><path stroke="#fff" stroke-linecap="round" d="M24 19.0005L24.0083 29"/></g>`),
		g.Group(children),
	)
}