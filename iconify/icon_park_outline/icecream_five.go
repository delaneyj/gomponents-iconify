package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcecreamFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M14 14c0-5.523 4.477-10 10-10s10 4.477 10 10v17.73a.27.27 0 0 1-.27.27H14.27a.27.27 0 0 1-.27-.27V14Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M21 16v6m6-6v6m-6 10v9a3 3 0 1 0 6 0v-9"/></g>`),
		g.Group(children),
	)
}