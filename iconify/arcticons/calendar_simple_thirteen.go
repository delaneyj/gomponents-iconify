package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleThirteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.932 20.925l3.5-1.925m0 0v14m9.179-7a3.467 3.467 0 0 1 3.457 3.457h0a3.467 3.467 0 0 1-3.457 3.457h-1.383c-2.42 0-3.284-.346-4.32-1.21m-.001-11.408c1.037-.864 2.074-1.21 4.321-1.21h1.383a3.467 3.467 0 0 1 3.457 3.457h0A3.467 3.467 0 0 1 27.61 26h-3.457"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}