package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CycleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 20c0-8 4-12 12-12m22 22c0 8-4 12-12 12m0-24c0-5.523 4.477-10 10-10h4v14H28v-4ZM6 28h14v4c0 5.523-4.477 10-10 10H6V28Z"/>`),
		g.Group(children),
	)
}