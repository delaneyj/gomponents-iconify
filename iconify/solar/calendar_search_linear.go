package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSearchLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 14v-2c0-3.771 0-5.657-1.172-6.828C19.657 4 17.771 4 14 4h-4C6.229 4 4.343 4 3.172 5.172C2 6.343 2 8.229 2 12v2c0 3.771 0 5.657 1.172 6.828C4.343 22 6.229 22 10 22h4M7 4V2.5M17 4V2.5"/><circle cx="18" cy="18" r="3"/><path stroke-linecap="round" d="M20.5 20.5L22 22M2.5 9h19"/></g>`),
		g.Group(children),
	)
}