package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M42 24L6 24"/><path fill="#2F88FF" d="M14 4L36 16H14V4Z"/><path fill="#2F88FF" d="M14 44V32H36L14 44Z"/></g>`),
		g.Group(children),
	)
}