package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotionSensorUrgentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.275 22L18 11.975L23.725 22h-11.45ZM18 21q.2 0 .35-.15t.15-.35q0-.2-.15-.35T18 20q-.2 0-.35.15t-.15.35q0 .2.15.35T18 21Zm-.5-2h1v-4h-1v4ZM2 22v-6h2v4h4v2H2ZM2 8V2h6v2H4v4H2Zm9 10.9q-2.3-.35-3.925-1.975T5.1 13h2q.3 1.475 1.363 2.537T11 16.9v2ZM5.1 11q.35-2.3 1.975-3.938T11 5.1v2q-1.475.3-2.538 1.363T7.1 11h-2Zm6.9 3q-.825 0-1.413-.588T10 12q0-.85.588-1.425T12 10q.85 0 1.425.575T14 12q0 .825-.575 1.413T12 14Zm4.9-3q-.3-1.475-1.363-2.538T13 7.1v-2q2.3.35 3.938 1.975T18.9 11h-2ZM20 8V4h-4V2h6v6h-2Z"/>`),
		g.Group(children),
	)
}