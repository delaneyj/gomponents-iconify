package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamonds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M12 8H36L44 18L24 42L4 18L12 8Z"/><path stroke="#fff" d="M4 18L44 18"/><path stroke="#fff" d="M24 42L16 18"/><path stroke="#fff" d="M24 42L32 18"/><path stroke="#000" d="M8 13L4 18L24 42L44 18L40 13"/></g>`),
		g.Group(children),
	)
}