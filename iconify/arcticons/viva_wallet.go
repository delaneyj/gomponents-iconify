package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VivaWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.674 10.081l9.826 4.704l-9.244 18.673c-3.93 7.939-8.082 3.754-10.643-1.77l-9.963-21.49L4.5 14.69l8.358 18.278c5.243 11.466 10.378-.47 10.755-1.28L33.674 10.08Z"/>`),
		g.Group(children),
	)
}