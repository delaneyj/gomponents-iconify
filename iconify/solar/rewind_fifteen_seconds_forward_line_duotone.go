package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindFifteenSecondsForwardLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M10 4.5L12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10c0-4.1-2.468-7.625-6-9.168" opacity=".5"/><path d="M16.5 8.5h-2.64a.5.5 0 0 0-.474.342l-.667 2a.5.5 0 0 0 .475.658H14.5a2 2 0 1 1 0 4h-2"/><path stroke-linejoin="round" d="m7.5 10.5l2.5-2v7"/></g>`),
		g.Group(children),
	)
}