package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrailLengthMediumOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17v-2h5q-.35-.425-.563-.925T11.1 13H5v-2h6.1q.125-.575.338-1.075T12 9H7V7h9q2.075 0 3.538 1.463T21 12q0 2.075-1.463 3.538T16 17H7Zm9-2q1.25 0 2.125-.875T19 12q0-1.25-.875-2.125T16 9q-1.25 0-2.125.875T13 12q0 1.25.875 2.125T16 15ZM3 17v-2h3v2H3Zm13-5Z"/>`),
		g.Group(children),
	)
}