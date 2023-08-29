package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbIncandescentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-3h2v3h-2Zm-9-9v-2h3v2H2Zm17 0v-2h3v2h-3Zm-1 6.9l-2.1-2.1l1.4-1.4l2.1 2.1l-1.4 1.4Zm-12 0l-1.4-1.4l2.1-2.1l1.4 1.4L6 19.9Zm6-2.9q-2.075 0-3.538-1.463T7 12q0-1.2.537-2.238T9 8V3h6v5q.925.725 1.463 1.763T17 12q0 2.075-1.463 3.538T12 17Zm-1-9.9q.25-.05.5-.075T12 7q.25 0 .5.025t.5.075V5h-2v2.1Zm1 7.9q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm0-3Z"/>`),
		g.Group(children),
	)
}