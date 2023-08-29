package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 21V3.5l.088-.059a11.56 11.56 0 0 1 12.824 0l.088.059V21M6 22.5h18m-15.5-6h-1a4 4 0 0 1-4-4V6m0 0c1 0 1.5-.5 1.5-.5V3M3.5 6C2.5 6 2 5.5 2 5.5V3m14 9c-.5 2-2 3-2 3v.5h2v.5c-.5 2-2 3-2 3M2 3h3M2 3V1m3 2V1m13.5 8.5h-7V5.043a11.559 11.559 0 0 1 7 0V9.5Z"/>`),
		g.Group(children),
	)
}