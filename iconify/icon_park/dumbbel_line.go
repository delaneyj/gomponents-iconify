package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DumbbelLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M13 10V38"/><path d="M44 24L4 24"/><path d="M35 10V38"/><path d="M7 14L7 34"/><path d="M41 14L41 34"/></g>`),
		g.Group(children),
	)
}