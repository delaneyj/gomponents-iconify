package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotionSensorAlertSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-6h2v4h4v2H2ZM2 6V0h6v2H4v4H2Zm9 10.9q-2.3-.35-3.925-1.975T5.1 11h2q.3 1.475 1.363 2.537T11 14.9v2ZM5.1 9q.35-2.3 1.975-3.938T11 3.1v2q-1.475.3-2.538 1.363T7.1 9h-2Zm6.9 3q-.825 0-1.413-.588T10 10q0-.85.588-1.425T12 8q.85 0 1.425.575T14 10q0 .825-.575 1.413T12 12Zm4.9-3q-.3-1.475-1.363-2.538T13 5.1v-2q2.3.35 3.938 1.975T18.9 9h-2ZM20 6V2h-4V0h6v6h-2Zm-1 16q-2.075 0-3.538-1.463T14 17q0-2.075 1.463-3.538T19 12q2.075 0 3.538 1.463T24 17q0 2.075-1.463 3.538T19 22Zm0-2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T19 19q-.2 0-.35.15t-.15.35q0 .2.15.35T19 20Zm-.5-2h1v-4h-1v4Z"/>`),
		g.Group(children),
	)
}