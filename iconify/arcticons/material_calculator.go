package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaterialCalculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.25" cy="33.25" r="11.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 27.18V7.45a2 2 0 0 0-2-1.95H24v21.68"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24H5.5v16.55a2 2 0 0 0 2 2h19.68M24 5.5H7.45a2 2 0 0 0-1.95 2V24H24Zm12.33 6.59l-6.16 6.15m0-6.15l6.16 6.15M14.37 37.6v-8.7m4.35 4.35h-8.71m27.59-1.14h-8.7m8.7 2.81h-8.7M18.7 15.15H10"/>`),
		g.Group(children),
	)
}