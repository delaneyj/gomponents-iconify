package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="10.5" cy="10.5" r="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 11.5v-5"/><circle cx="10.5" cy="14.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}