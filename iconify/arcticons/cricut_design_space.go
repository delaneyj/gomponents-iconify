package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CricutDesignSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.1 25.619c-.263.525-.963.962-1.575.962h0A1.942 1.942 0 0 1 8.6 24.656v-1.225c0-1.05.875-1.925 1.925-1.925h0c.7 0 1.313.35 1.575.963m14.612 3.15c-.262.525-.962.962-1.575.962h0a1.942 1.942 0 0 1-1.925-1.925v-1.225c0-1.05.875-1.925 1.925-1.925h0c.7 0 1.313.35 1.575.963"/><circle cx="19.887" cy="19.231" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.887 21.506v5.075m-4.725-3.237c0-1.05.875-1.925 1.925-1.925h0m-1.925.087v5.075m22.575-6.65v5.688c0 .525.35.962.963.962h.262m-2.274-5.162h2.1m-9.188.087v3.15c0 1.05.875 1.925 1.925 1.925h0c1.05 0 1.925-.875 1.925-1.925v-3.15m0 3.238v1.837"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}