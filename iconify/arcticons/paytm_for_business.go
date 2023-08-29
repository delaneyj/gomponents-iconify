package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaytmForBusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 32.689V7.5a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1m0 0v20.9a2 2 0 0 0 2 2h25.305M5.5 32.5h29.284M5.5 29.467h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.557 23.533v-7.991l4 8l4-7.988v7.988m-31.424 0v-8h2.619a2.687 2.687 0 0 1 0 5.374h-2.62m16.516-5.374h5.3m-2.65 8v-8m-2.65 6.216v2.7a2 2 0 0 1-2 2h0c-.53 0-1.04-.21-1.414-.586"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.648 18.458v3.3a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-3.3m-1.906 3.064a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2m0 3.3v-5.3"/><circle cx="38.5" cy="38.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.529 38.433a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm0 0h-3.301"/>`),
		g.Group(children),
	)
}