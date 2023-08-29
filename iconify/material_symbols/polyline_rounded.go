package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolylineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22q-.825 0-1.413-.588T15 20v-.5l-7.275-3.625l-.35.088Q7.225 16 7 16H5q-.825 0-1.412-.588T3 14v-2q0-.825.588-1.413T5 10h2.138q.062 0 .137.025l2.875-3.3q-.075-.15-.113-.338T10 6V4q0-.825.588-1.413T12 2h2q.825 0 1.413.588T16 4v2q0 .825-.588 1.413T14 8h-2.163q-.062 0-.112-.025l-2.85 3.275q.05.2.088.363T9 12v2.113q0 .062-.025.112l6.15 3.075q.2-.55.7-.925T17 16h2q.825 0 1.413.588T21 18v2q0 .825-.588 1.413T19 22h-2Z"/>`),
		g.Group(children),
	)
}