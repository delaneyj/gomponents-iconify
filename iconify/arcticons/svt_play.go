package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SvtPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.17 21.12l-3.16 8.36l-3.15-8.36M9.7 28.78a3.57 3.57 0 0 0 2.59.7H13a2.08 2.08 0 0 0 2.09-2.09h0A2.09 2.09 0 0 0 13 25.3h-1.41a2.08 2.08 0 0 1-2.09-2.09h0a2.09 2.09 0 0 1 2.09-2.09h.7a3.53 3.53 0 0 1 2.6.71m12.29-3.31v9.39a1.57 1.57 0 0 0 1.57 1.57h.48m-2.05-8.36h1.65m9.67 4.18l-6.25-4.18v8.36l6.25-4.18z"/>`),
		g.Group(children),
	)
}