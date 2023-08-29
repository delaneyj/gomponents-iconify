package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirtyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M4 8h40v36H4z"/><path stroke-linecap="round" d="M28 20v14h8V20h-8Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M17 4v8m14-8v8m-19 8h8v14h-8m8-7h-6"/></g>`),
		g.Group(children),
	)
}