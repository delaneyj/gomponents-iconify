package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.077 28.333c0 2.593 2.074 4.667 4.667 4.667s4.494-2.074 4.494-4.667v-4.666c0-2.593-2.075-4.667-4.494-4.667s-4.667 2.074-4.667 4.667v4.666Zm-5.611-2.419a3.467 3.467 0 0 1 3.457 3.456h0a3.467 3.467 0 0 1-3.457 3.457h-1.383c-2.42 0-3.283-.346-4.32-1.21m-.001-11.407C14.8 19.346 15.837 19 18.084 19h1.383a3.467 3.467 0 0 1 3.457 3.457h0a3.467 3.467 0 0 1-3.457 3.457H16.01"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}