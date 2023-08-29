package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneDm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.78 37.272h23.98c13.018 0 10.842-19.588-2.216-15.235c0-10.882-19.588-10.882-19.588 2.176C2.074 22.037 2.074 37.272 10.78 37.272Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.273 27.477L24 30.75l-3.273-3.273M24 30.75v-9.998m-5.758 12h11.516m8.915-15.223v-6.801m-3.401 3.401h6.801"/>`),
		g.Group(children),
	)
}