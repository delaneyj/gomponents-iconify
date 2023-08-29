package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tui(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.24 20.413v4.426a2.682 2.682 0 0 0 2.682 2.683h0a2.682 2.682 0 0 0 2.683-2.683v-4.426m0 4.426v2.682"/><circle cx="38.42" cy="17.127" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.42 20.413v7.108M20.646 18.2v7.98a1.341 1.341 0 0 0 1.341 1.341h.403m-3.152-7.108h2.816m-8.331 6.237l4.024-2.683l-4.024-2.682M8.359 26.65l4.023-2.683l-4.023-2.682m16.379 2.682h2.682"/>`),
		g.Group(children),
	)
}