package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieQuarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.519 2.747a8 8 0 1 0 9.705 9.845"/><path d="M18.5 10.5a8 8 0 0 0-8-8v8z"/></g>`),
		g.Group(children),
	)
}