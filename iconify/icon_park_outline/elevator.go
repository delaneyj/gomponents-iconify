package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M42 41V7H6v34h36Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 7v34m0-34H4h2v34M42 7h2m-2 34h2m-2 0H6m0 0H4M24 7v34m9-21v8m-3-5l3-3l3 3m-21 5v-8"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m12 25l3 3l3-3"/></g>`),
		g.Group(children),
	)
}