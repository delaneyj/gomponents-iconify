package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StandUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M17 22.458c-4.057 1.274-7 5.065-7 9.542c0 5.523 4.477 10 10 10a9.985 9.985 0 0 0 8.662-5"/><path stroke-linecap="round" stroke-linejoin="round" d="m17 18l1.5-2H30l-7 14h14v14"/><circle cx="34" cy="8" r="4"/></g>`),
		g.Group(children),
	)
}