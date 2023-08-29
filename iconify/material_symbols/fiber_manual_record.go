package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberManualRecord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q-2.925 0-4.963-2.038T5 12q0-2.925 2.038-4.963T12 5q2.925 0 4.963 2.038T19 12q0 2.925-2.038 4.963T12 19Z"/>`),
		g.Group(children),
	)
}