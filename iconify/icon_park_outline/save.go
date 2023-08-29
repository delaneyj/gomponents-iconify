package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M6 9a3 3 0 0 1 3-3h25.281L42 13.207V39a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Z"/><path d="M24.008 6L24 13.385c0 .34-.448.615-1 .615h-8c-.552 0-1-.275-1-.615V6" clip-rule="evenodd"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M24.008 6L24 13.385c0 .34-.448.615-1 .615h-8c-.552 0-1-.275-1-.615V6h10.008Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 6h25.281M14 26h20m-20 8h10.008"/></g>`),
		g.Group(children),
	)
}