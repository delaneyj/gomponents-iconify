package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RocketOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 19.05l1.975-.8q-.25-.725-.463-1.475t-.337-1.5l-1.175.8v2.975ZM10 18h4q.45-1 .725-2.438T15 12.626q0-2.475-.825-4.688T12 4.526q-1.35 1.2-2.175 3.413T9 12.625q0 1.5.275 2.938T10 18Zm2-5q-.825 0-1.413-.588T10 11q0-.825.588-1.413T12 9q.825 0 1.413.588T14 11q0 .825-.588 1.413T12 13Zm6 6.05v-2.975l-1.175-.8q-.125.75-.338 1.5t-.462 1.475l1.975.8ZM12 1.975q2.475 1.8 3.738 4.575T17 13l2.1 1.4q.425.275.663.725t.237.95V22l-4.975-2h-6.05L4 22v-5.925q0-.5.238-.95T4.9 14.4L7 13q0-3.675 1.263-6.45T12 1.975Z"/>`),
		g.Group(children),
	)
}