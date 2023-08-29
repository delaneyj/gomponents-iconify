package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Esea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.054 2.767L8.95 9.511L0 9.65l5.832 4.47L1.042 21l8.491-4.088l5.711 4.322V14.12L24 9.796l-17.255 4.02a12.575 12.575 0 0 0 1.589-1.955a5.475 5.475 0 0 0 .617-1.786l5.593-.15z"/>`),
		g.Group(children),
	)
}