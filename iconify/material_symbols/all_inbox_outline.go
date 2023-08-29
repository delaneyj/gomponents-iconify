package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllInboxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16h12v-3h-2.55q-.525.925-1.45 1.463T14 15q-1.05 0-1.975-.537T10.55 13H8v3Zm6-3q.85 0 1.425-.588T16 11h4V4H8v7h4q0 .825.588 1.413T14 13Zm-6 5q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V6h2v14h14v2H4Zm4-6h12H8Z"/>`),
		g.Group(children),
	)
}