package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.077 28.333c0 2.593 2.074 4.667 4.667 4.667s4.494-2.074 4.494-4.667v-4.666c0-2.593-2.075-4.667-4.494-4.667s-4.667 2.074-4.667 4.667v4.666Zm-11.315-4.666c0-2.593 2.075-4.667 4.494-4.667s4.667 2.074 4.667 4.667c0 1.21-.518 2.42-1.383 3.284c-1.9 1.555-7.778 6.05-7.778 6.05h9.161"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}