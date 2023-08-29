package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.75 14.95L16.3 13.5q.35-.575.525-1.2T17 11h2q0 1.1-.325 2.087t-.925 1.863Zm-2.95-3L9 6.15V5q0-1.25.875-2.125T12 2q1.25 0 2.125.875T15 5v6q0 .275-.063.5t-.137.45ZM11 21v-3.1q-2.6-.35-4.3-2.312T5 11h2q0 2.075 1.463 3.538T12 16q.85 0 1.613-.263T15 15l1.425 1.425q-.725.575-1.588.963T13 17.9V21h-2Zm8.8 1.6L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}