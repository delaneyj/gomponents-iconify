package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideReddit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 15.58V7.45a2 2 0 0 0-2-1.95H7.45a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h8.21a19.2 19.2 0 0 1-.35-5.36a14.73 14.73 0 0 1 6.32-10.82A22.21 22.21 0 0 1 25.75 24h0A10.48 10.48 0 1 1 27 28.9a15 15 0 0 0-2.46 1.42a9.38 9.38 0 0 0-4.28 7.24a15.14 15.14 0 0 0 .51 4.94h19.78a2 2 0 0 0 2-2v-8.08"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27 28.9a10.48 10.48 0 0 1-1.25-4.9"/>`),
		g.Group(children),
	)
}