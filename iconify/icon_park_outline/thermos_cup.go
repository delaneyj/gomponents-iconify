package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermosCup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M10 31s.071 6 14 6s14-6 14-6V15H10v16Z"/><path stroke-linecap="round" d="M24 4v6m-8-5v4m16-4v4M14 36v5a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5"/></g>`),
		g.Group(children),
	)
}