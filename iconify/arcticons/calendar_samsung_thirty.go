package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSamsungThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.077 26.333c0 2.593 2.074 4.667 4.667 4.667s4.494-2.074 4.494-4.667v-4.666c0-2.593-2.075-4.667-4.494-4.667s-4.667 2.074-4.667 4.667v4.666Zm-5.611-2.419a3.467 3.467 0 0 1 3.457 3.456h0a3.467 3.467 0 0 1-3.457 3.457h-1.383c-2.42 0-3.283-.346-4.32-1.21m-.001-11.407C14.8 17.346 15.837 17 18.084 17h1.383a3.467 3.467 0 0 1 3.457 3.457h0a3.467 3.467 0 0 1-3.457 3.457H16.01"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-4v1a3.5 3.5 0 1 1-7 0v-1h-8v1a3.5 3.5 0 1 1-7 0v-1h-3a4 4 0 0 0-4 4v29a4 4 0 0 0 4 4h29a4 4 0 0 0 4-4v-29a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}