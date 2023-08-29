package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserSegmentTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.318 2.44a1.5 1.5 0 0 0-2.121 0l-8.76 8.759a1.5 1.5 0 0 0 0 2.121l4.243 4.243c.313.313.73.459 1.14.437h6.265a1.5 1.5 0 1 0 0-1h-4.72l8.196-8.197a1.5 1.5 0 0 0 0-2.121l-4.243-4.243Zm-8.487 7.78l4.95 4.949l-1.687 1.687a.5.5 0 0 1-.707 0l-4.243-4.243a.5.5 0 0 1 0-.707l1.687-1.687Z"/>`),
		g.Group(children),
	)
}