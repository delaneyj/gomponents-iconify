package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Level(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 42L4 18.5L9.69488 6L38.3051 6L44 18.5L24 42Z"/><path stroke="#fff" d="M32 18L24 27L16 18"/></g>`),
		g.Group(children),
	)
}