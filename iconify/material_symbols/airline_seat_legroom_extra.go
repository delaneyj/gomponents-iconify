package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomExtra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17H4q-.825 0-1.413-.588T2 15V3h2v12h9v2Zm5.4 4L15 14H8.5q-1.25 0-2.125-.875T5.5 11V3h6v6h3q.575 0 1.05.313t.75.837l3.4 6.95l1.1-.5q.575-.275 1.163-.088t.887.738q.3.575.088 1.175t-.788.875L18.4 21Z"/>`),
		g.Group(children),
	)
}