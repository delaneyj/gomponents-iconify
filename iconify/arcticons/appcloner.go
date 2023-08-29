package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appcloner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 33.91H26a8.36 8.36 0 0 0 7.88-7.85v-12m0-2V5.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 14.09H22a8.36 8.36 0 0 0-7.88 7.85v12m0 2v6.56m2.48-25.9l14.76 14.76M14.09 24L24 33.91m0-19.82L33.91 24"/>`),
		g.Group(children),
	)
}