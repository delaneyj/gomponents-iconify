package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CounterClockwiseClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.15 7.5c0-2.835-2.21-5.65-5.65-5.65c-2.778 0-4.152 2.056-4.737 3.15H4.5a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v1.813C2.705 3.071 4.334.85 7.5.85c4.063 0 6.65 3.335 6.65 6.65c0 3.315-2.587 6.65-6.65 6.65c-1.944 0-3.562-.77-4.715-1.942a6.772 6.772 0 0 1-1.427-2.167a.5.5 0 1 1 .925-.38c.28.681.692 1.314 1.216 1.846c.972.99 2.336 1.643 4.001 1.643c3.44 0 5.65-2.815 5.65-5.65ZM7.5 4a.5.5 0 0 1 .5.5v2.793l1.854 1.853a.5.5 0 0 1-.708.708l-2-2A.5.5 0 0 1 7 7.5v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}