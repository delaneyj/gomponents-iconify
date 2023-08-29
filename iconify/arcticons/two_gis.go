package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoGis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.578 36.142c-.04-4.789 1.063-8.746 7.082-8.498c8.586-21.881-26.424-21.047-17.092-.13c2.83-.176 7.521 1.509 6.72 9.146M42.5 33.479l-37 5.824m0-28.263l12.345 2.29m14.51 2.2L42.5 17.011"/>`),
		g.Group(children),
	)
}