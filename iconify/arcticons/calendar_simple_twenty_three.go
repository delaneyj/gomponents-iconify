package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleTwentyThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.762 23.667c0-2.593 2.075-4.667 4.494-4.667s4.667 2.074 4.667 4.667c0 1.21-.518 2.42-1.383 3.284c-1.9 1.555-7.778 6.05-7.778 6.05h9.161m7.857-7.087a3.467 3.467 0 0 1 3.457 3.456h0a3.467 3.467 0 0 1-3.456 3.457h-1.383c-2.42 0-3.284-.346-4.321-1.21m0-11.407c1.037-.864 2.074-1.21 4.321-1.21h1.383a3.467 3.467 0 0 1 3.457 3.457h0a3.467 3.467 0 0 1-3.457 3.457h-3.457"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}