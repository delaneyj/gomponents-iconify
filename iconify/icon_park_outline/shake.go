package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="m4 10l4 4.667l-4 4.666L8 24l-4 4.667l4 4.666L4 38m40-28l-4 4.667l4 4.666L40 24l4 4.667l-4 4.666L44 38"/><path d="M34 6H14v36h20V6Z"/><path stroke-linecap="round" d="M22 35h4"/></g>`),
		g.Group(children),
	)
}