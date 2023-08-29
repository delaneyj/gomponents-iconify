package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvePortal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.333 15.214H4.5v1.084m0 2.979h10.833M4.5 22.256v1.084h10.833m3.521-8.126L24 24.423l5.146-9.208m14.354-.056H32.667v1.083m0 3.035H43.5M32.667 22.2v1.084H43.5m-38.325 9.51v-3.8H6.41a1.282 1.282 0 1 1 0 2.565H5.175m7.088 0a1.259 1.259 0 1 0 2.518 0v-1.283a1.277 1.277 0 0 0-1.283-1.282a1.238 1.238 0 0 0-1.235 1.282Zm7.089 1.235v-3.8h1.235a1.282 1.282 0 1 1 0 2.565h-1.235m1.282-.048l1.235 1.283m4.571-3.8h2.517m-1.235 3.8v-3.8m8.323 3.8l-1.235-3.8l-1.282 3.8m7.397-3.8v3.8h1.9"/>`),
		g.Group(children),
	)
}