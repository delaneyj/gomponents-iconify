package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSimpleEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.137 26h-2.275a3.51 3.51 0 0 0-3.5 3.5h0a3.51 3.51 0 0 0 3.5 3.5h2.276a3.51 3.51 0 0 0 3.5-3.5h0a3.51 3.51 0 0 0-3.5-3.5Zm0 0a3.51 3.51 0 0 0 3.5-3.5h0a3.51 3.51 0 0 0-3.5-3.5h-2.275a3.51 3.51 0 0 0-3.5 3.5h0a3.51 3.51 0 0 0 3.5 3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.099 7.859c-1.06 0-1.92.86-1.92 1.92V41.58c0 1.06.86 1.919 1.92 1.919H39.9c1.06 0 1.92-.86 1.92-1.92h0v-31.8c0-1.06-.86-1.92-1.92-1.92H8.1Zm4.718 2.399V4.5m22.366 5.758V4.5"/>`),
		g.Group(children),
	)
}