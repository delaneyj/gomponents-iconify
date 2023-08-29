package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkPingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18q-.425 0-.713-.288T4 17q0-.425.288-.713T5 16h5.5L2.7 8.2q-.275-.275-.288-.687T2.7 6.8q.275-.275.7-.275t.7.275l7.9 7.875l5.2-5.2q-.1-.225-.15-.462T17 8.5q0-1.05.725-1.775T19.5 6q1.05 0 1.775.725T22 8.5q0 1.05-.725 1.775T19.5 11q-.225 0-.438-.037t-.412-.113L13.5 16H19q.425 0 .713.288T20 17q0 .425-.288.713T19 18H5Z"/>`),
		g.Group(children),
	)
}