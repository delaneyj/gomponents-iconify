package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cyrillic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.265 26.875v-8l5.3 8v-8m13.531-.015v8m5.3-8v8m-5.3-4.015h5.3m-15.666-1.27v6a1.99 1.99 0 0 1-1.98 2h0a1.962 1.962 0 0 1-.648-.11"/><circle cx="18.731" cy="19.175" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.396 23.56a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v1.3a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2M25.5 18.925l-3 10.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}