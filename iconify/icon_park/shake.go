package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M4 10L8 14.6667L4 19.3333L8 24L4 28.6667L8 33.3333L4 38"/><path stroke="#000" stroke-linecap="round" d="M44 10L40 14.6667L44 19.3333L40 24L44 28.6667L40 33.3333L44 38"/><path fill="#2F88FF" stroke="#000" d="M34 6H14V42H34V6Z"/><path stroke="#fff" stroke-linecap="round" d="M22 35H26"/></g>`),
		g.Group(children),
	)
}