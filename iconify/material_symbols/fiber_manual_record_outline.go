package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberManualRecordOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12Zm0 7q-2.9 0-4.95-2.05T5 12q0-2.9 2.05-4.95T12 5q2.9 0 4.95 2.05T19 12q0 2.9-2.05 4.95T12 19Zm0-2q2.075 0 3.538-1.463T17 12q0-2.075-1.463-3.538T12 7Q9.925 7 8.462 8.463T7 12q0 2.075 1.463 3.538T12 17Z"/>`),
		g.Group(children),
	)
}