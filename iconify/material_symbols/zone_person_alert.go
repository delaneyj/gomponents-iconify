package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZonePersonAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18v-3h2v3h3v2H4Zm4.8-4l1-5.1l-1.8.7V15H6v-4.7l3.95-1.7q.875-.375 1.288-.487T12 8q.525 0 .975.275T13.7 9l1 1.6l.2.3q.1.15.225.275q-1.225.8-2.025 2.063T12.075 16H8.8Zm4.7-8.5q-.825 0-1.413-.588T11.5 5.5q0-.825.588-1.413T13.5 3.5q.825 0 1.413.588T15.5 5.5q0 .825-.588 1.413T13.5 7.5ZM2 5V2q0-.825.588-1.413T4 0h3v2H4v3H2Zm18 0V2h-3V0h3q.825 0 1.413.588T22 2v3h-2Zm-1 17q-2.075 0-3.538-1.463T14 17q0-2.075 1.463-3.538T19 12q2.075 0 3.538 1.463T24 17q0 2.075-1.463 3.538T19 22Zm-.5-4h1v-4h-1v4Zm.5 2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T19 19q-.2 0-.35.15t-.15.35q0 .2.15.35T19 20Z"/>`),
		g.Group(children),
	)
}