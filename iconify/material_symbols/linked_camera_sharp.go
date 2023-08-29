package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedCameraSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.65 7q0-1.95-1.35-3.3T16 2.35V1q2.5 0 4.25 1.75T22 7h-1.35ZM18 7q0-.825-.588-1.413T16 5V3.65q1.375 0 2.337.975T19.35 7H18Zm-6 10.5q1.875 0 3.188-1.313T16.5 13q0-1.875-1.313-3.188T12 8.5q-1.875 0-3.188 1.313T7.5 13q0 1.875 1.313 3.188T12 17.5Zm0-2q-1.05 0-1.775-.725T9.5 13q0-1.05.725-1.775T12 10.5q1.05 0 1.775.725T14.5 13q0 1.05-.725 1.775T12 15.5ZM2 21V5h5.15L9 3h6v3q.825 0 1.413.588T17 8h5v13H2Z"/>`),
		g.Group(children),
	)
}