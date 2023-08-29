package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KukuFm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.711 27.038v-6.076l3.096 6.076l3.096-6.076v6.076m-10.981-6.076h3.096M29.922 24h2.013m-2.013-3.038v6.076M4.5 20.962v6.076m3.328 0L5.274 24l2.554-3.038M5.274 24H4.5m16.22-3.038v4.026a2.013 2.013 0 1 0 4.025 0v-4.026m-9.2 0v6.076m3.329 0L16.319 24l2.555-3.038M16.319 24h-.774m-5.871-3.038v4.026a2.013 2.013 0 1 0 4.025 0v-4.026"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.322 18.517v10.967H43.5V18.515Z"/>`),
		g.Group(children),
	)
}