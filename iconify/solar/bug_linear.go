package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 15v-3.063A3.938 3.938 0 0 0 15.062 8H8.939A3.938 3.938 0 0 0 5 11.938V15a7 7 0 1 0 14 0Z"/><path d="M16.5 8.5v-1a4.5 4.5 0 1 0-9 0v1"/><path stroke-linecap="round" d="M19 14h3M5 14H2M14.5 3.5L17 2M9.5 3.5L7 2m13.5 18l-2-.8m2-11.2l-2 .8M3.5 20l2-.8M3.5 8l2 .8M12 21.5V15"/></g>`),
		g.Group(children),
	)
}