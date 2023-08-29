package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarDateLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12c0-3.771 0-5.657 1.172-6.828C4.343 4 6.229 4 10 4h4c3.771 0 5.657 0 6.828 1.172C22 6.343 22 8.229 22 12v2c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14v-2Z"/><path stroke-linecap="round" d="M7 4V2.5M17 4V2.5M2 9h20" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 14.5l1.5-1.5v4"/><path stroke-linecap="round" d="M13 16v-2a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0Z"/></g>`),
		g.Group(children),
	)
}