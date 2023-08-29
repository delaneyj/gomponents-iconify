package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSamsungTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.077 26.333c0 2.593 2.074 4.667 4.667 4.667s4.494-2.074 4.494-4.667v-4.666c0-2.593-2.075-4.667-4.494-4.667s-4.667 2.074-4.667 4.667v4.666Zm-11.315-4.666c0-2.593 2.075-4.667 4.494-4.667s4.667 2.074 4.667 4.667c0 1.21-.518 2.42-1.383 3.284c-1.9 1.555-7.778 6.05-7.778 6.05h9.161"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-4v1a3.5 3.5 0 1 1-7 0v-1h-8v1a3.5 3.5 0 1 1-7 0v-1h-3a4 4 0 0 0-4 4v29a4 4 0 0 0 4 4h29a4 4 0 0 0 4-4v-29a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}