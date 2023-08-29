package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSamsungTwentyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.734 21.667c0-2.593 2.074-4.667 4.494-4.667s4.666 2.074 4.666 4.667c0 1.21-.518 2.42-1.382 3.284C19.61 26.506 13.734 31 13.734 31h9.16m2.097-1.225c1.05.875 1.925 1.225 4.2 1.225h.525c2.45 0 4.55-2.1 4.55-4.55h0c0-2.45-2.1-4.55-4.55-4.55h-4.725V17h9.275"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-4v1a3.5 3.5 0 1 1-7 0v-1h-8v1a3.5 3.5 0 1 1-7 0v-1h-3a4 4 0 0 0-4 4v29a4 4 0 0 0 4 4h29a4 4 0 0 0 4-4v-29a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}