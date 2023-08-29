package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 14.5v-2a3 3 0 0 1 3-3h8"/><path d="m11.499 12.5l3.001-3l-3.001-3"/><path d="m14.499 12.5l3.001-3l-3.001-3"/></g>`),
		g.Group(children),
	)
}