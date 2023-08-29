package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="38" height="24" x="5" y="18" stroke-linecap="round" rx="2"/><path stroke-linecap="round" d="M8 12h32M15 6h18m-7 18v6"/><path d="M18 32.75c0-1.52 1.29-2.75 2.88-2.75H26v3.25c0 1.52-1.29 2.75-2.88 2.75h-2.24C19.29 36 18 34.77 18 33.25v-.5Z"/><path stroke-linecap="round" d="m31 25l-5-1"/></g>`),
		g.Group(children),
	)
}