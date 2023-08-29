package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArianeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.732 24.925a2 2 0 0 1 2-2h0m-2 0v5.3"/><circle cx="20.537" cy="20.475" r=".7" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.537 22.925v5.3m17.5-1.009a2 2 0 0 1-1.737 1.009h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.65h-4m-7.92.65a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2m0 3.3v-5.3m5.98 5.3v-3.3a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2v3.3m.001-3.3v-2M9.7 28.201l2.599-7.976m2.601 8l-2.601-8m1.731 5.324h-3.465"/>`),
		g.Group(children),
	)
}