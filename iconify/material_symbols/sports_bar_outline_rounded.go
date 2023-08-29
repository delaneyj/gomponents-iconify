package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsBarOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19h7V9h-2q-.7 0-1.15.35t-1.075 1.025q-.5.55-1.163 1.137T8 12.375V19Zm-2 0v-6.65q-1.3-.35-2.15-1.4T3 8.5q0-1.325.763-2.35t1.962-1.425q.575-1.2 1.713-1.95T10 2.025q.875 0 1.638.3t1.387.8q.25-.05.475-.088T14 3q1.65 0 2.825 1.175T18 7q0 .55-.138 1.05T17.45 9H19q.825 0 1.413.588T21 11v6q0 .825-.588 1.413T19 19h-2q0 .825-.588 1.413T15 21H8q-.825 0-1.413-.588T6 19ZM5 8.5q0 .825.588 1.413T7 10.5q.8 0 1.363-.525T9.524 8.8q.625-.675 1.413-1.238T13 7h3q0-.825-.588-1.413T14 5q-.625 0-1.05.163l-.425.162l-.775-.65q-.275-.225-.713-.438T10 4.025q-.8 0-1.463.425T7.525 5.6l-.35.75l-.8.275q-.625.2-1 .713T5 8.5ZM17 17h2v-6h-2v6Zm-9 2h7h-7Z"/>`),
		g.Group(children),
	)
}