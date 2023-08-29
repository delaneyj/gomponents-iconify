package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlyDelta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.363 41.028l-6.841-19.55L40.5 7.5l-7.137 33.528ZM31.989 44.5l-8.23-20.258l-6.015 6.014L31.99 44.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.972 14.637L40.5 7.5L26.522 21.478l-19.55-6.841M3.5 16.011l14.244 14.245l6.014-6.014L3.5 16.011"/>`),
		g.Group(children),
	)
}