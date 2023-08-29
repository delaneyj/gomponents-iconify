package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM9.9 31.2V16.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.9 16.8h6a3.59 3.59 0 0 1 3.6 3.6h0a3.59 3.59 0 0 1-3.6 3.6h-6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.9 24h6a3.59 3.59 0 0 1 3.6 3.6h0a3.59 3.59 0 0 1-3.6 3.6h-6m15.5 0h0a3.59 3.59 0 0 1-3.6-3.6v-2.2a3.59 3.59 0 0 1 3.6-3.6h0a3.59 3.59 0 0 1 3.6 3.6v2.2a3.59 3.59 0 0 1-3.6 3.6Zm10-12.4v10.6a1.79 1.79 0 0 0 1.8 1.8h.7m-6.6-12.4v12.4m2.2-9.4h1.9m0 0h1.9"/><circle cx="25.4" cy="33.3" r=".8" fill="currentColor"/>`),
		g.Group(children),
	)
}