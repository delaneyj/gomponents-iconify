package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataArrival(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 42h36V6H30c-1.324 3.159-3.324 4.738-6 4.738S19.324 9.158 18 6H6v36Z" clip-rule="evenodd"/><path stroke-linecap="round" d="m15 24l6 6l12-12"/></g>`),
		g.Group(children),
	)
}