package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.962 24h14.076m-16.307 8.385v-16.77m-3.962 15.47v-14.17m-3.961 12.87v-11.57M3.5 24h1.038m28.731-8.385v16.77m3.962-15.47v14.17m3.961-12.87v11.57M44.5 24h-1.038"/>`),
		g.Group(children),
	)
}