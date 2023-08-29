package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pesticide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M15 11.368V4h18v7.368l6 6.119V42a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2V17.486l6-6.118Z" clip-rule="evenodd"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 23h8v12H9"/><path stroke-linecap="round" d="M15 11.5h18M31 23v6m0 5v1"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 38V20"/></g>`),
		g.Group(children),
	)
}