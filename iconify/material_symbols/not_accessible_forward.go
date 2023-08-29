package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotAccessibleForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L18.15 21H17v-1.15L13.15 16H11q-1.1 0-1.7-.938T9.15 13.1l.35-.75l-8.125-8.125L2.8 2.8l18.4 18.4l-1.425 1.425ZM19 16.15l-2.65-2.65H17q.825 0 1.413.588T19 15.5v.65Zm-4.175-4.175L9.85 7h4.1q1.125 0 1.713.913T15.8 9.85l-.975 2.125ZM8 22q-2.075 0-3.538-1.463T3 17q0-2.075 1.463-3.538T8 12v2q-1.25 0-2.125.875T5 17q0 1.25.875 2.125T8 20q1.25 0 2.125-.875T11 17h2q0 2.075-1.463 3.538T8 22Zm8-15.5q-.825 0-1.413-.588T14 4.5q0-.825.588-1.413T16 2.5q.825 0 1.413.588T18 4.5q0 .825-.588 1.413T16 6.5Z"/>`),
		g.Group(children),
	)
}