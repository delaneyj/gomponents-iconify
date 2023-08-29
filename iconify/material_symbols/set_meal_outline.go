package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetMealOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15q-.825 0-1.413-.588T1 13V4q0-.825.588-1.413T3 2h18q.825 0 1.413.588T23 4v9q0 .825-.588 1.413T21 15H3ZM3 4v9h18V4H3Zm.075 14.5L3 17l17.975-.95l.075 1.5l-17.975.95ZM3 20.975v-1.5h18v1.5H3Zm7.25-9.475q1.85 0 3.563-.65t2.987-2q.15 1.05 1.1 1.6T20 11V6q-1.15 0-2.1.563T16.8 8.2q-1.325-1.3-3.013-2t-3.537-.7q-1.975 0-3.8.688T3.5 8.5q1.125 1.625 2.95 2.313t3.8.687ZM3 4v9v-9Z"/>`),
		g.Group(children),
	)
}