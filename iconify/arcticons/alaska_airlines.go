package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlaskaAirlines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.876 37.235c6.053-11.607 15.906-22.837 20.624-26.47c-3.215 5.929-2.63 23.63-2.13 26.47m-3.701-14.237H11.5"/>`),
		g.Group(children),
	)
}