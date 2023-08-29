package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrailLengthShortOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17q-1.825 0-3.188-1.137T10.1 13H4v-2h6.1q.125-.575.338-1.075T11 9H6V7h9q2.075 0 3.538 1.463T20 12q0 2.075-1.463 3.538T15 17Zm0-2q1.25 0 2.125-.875T18 12q0-1.25-.875-2.125T15 9q-1.25 0-2.125.875T12 12q0 1.25.875 2.125T15 15Zm-8 2v-2h3v2H7Zm8-5Z"/>`),
		g.Group(children),
	)
}