package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatchPrediction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-.825 0-1.413-.588T5 20V10q0-.825.588-1.413T7 8h10q.825 0 1.413.588T19 10v10q0 .825-.588 1.413T17 22H7Zm4-1.5h2V19h-2v1.5Zm0-2.5h2q0-.575.388-1.15t.862-1.175q.475-.625.863-1.275T15.5 13q0-1.45-1.025-2.475T12 9.5q-1.45 0-2.475 1.025T8.5 13q0 .75.388 1.4t.862 1.275q.475.6.863 1.175T11 18ZM6 6.5q0-.625.438-1.063T7.5 5h9q.625 0 1.063.438T18 6.5H6Zm1-3q0-.625.438-1.063T8.5 2h7q.625 0 1.063.438T17 3.5H7Z"/>`),
		g.Group(children),
	)
}