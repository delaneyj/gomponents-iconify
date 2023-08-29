package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatchPredictionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22V8h14v14H5Zm6-1.5h2V19h-2v1.5Zm0-2.5h2q0-.575.388-1.15t.862-1.175q.475-.625.863-1.275T15.5 13q0-1.45-1.025-2.475T12 9.5q-1.45 0-2.475 1.025T8.5 13q0 .75.388 1.4t.862 1.275q.475.6.863 1.175T11 18ZM6 6.5V5h12v1.5H6Zm1-3V2h10v1.5H7Z"/>`),
		g.Group(children),
	)
}