package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleTwentyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.734 23.667c0-2.593 2.074-4.667 4.494-4.667s4.666 2.074 4.666 4.667c0 1.21-.518 2.42-1.382 3.284C19.61 28.506 13.734 33 13.734 33h9.16m2.097-1.225c1.05.875 1.925 1.225 4.2 1.225h.525c2.45 0 4.55-2.1 4.55-4.55h0c0-2.45-2.1-4.55-4.55-4.55h-4.725V19h9.275"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}