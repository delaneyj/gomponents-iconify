package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SparssDecsync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.09 10.55A16.24 16.24 0 0 0 7.76 24m7.16 13.46A16.23 16.23 0 0 0 40.23 24M9.04 24h2.82l-3.88 5.92L4.11 24h4.93zm30.14 0h-2.82l3.87-5.92L44.11 24h0h-4.93z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.528 28.99A13.518 13.518 0 0 0 19.01 15.472M26.8 28.99a7.79 7.79 0 0 0-7.79-7.79"/><circle cx="19.01" cy="28.99" r="2.482" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}