package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20q-.425 0-.713-.288T19 19v-1q0-.425.288-.713T20 17h2v-1h-2.5q-.2 0-.35-.15T19 15.5q0-.2.15-.35t.35-.15H22q.425 0 .713.288T23 16v1q0 .425-.288.713T22 18h-2v1h2.5q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15H20ZM7.925 18q-.675 0-.988-.575t.038-1.15l3.525-5.55L7.3 5.7q-.35-.55-.038-1.125T8.226 4q.3 0 .55.138t.4.387L11.95 9h.1l2.75-4.475q.15-.25.413-.387T15.75 4q.675 0 .975.575t-.05 1.15l-3.2 5l3.55 5.55q.35.575.025 1.15t-.975.575q-.3 0-.55-.138t-.4-.387l-3.075-4.9h-.1l-3.075 4.9q-.15.25-.413.388T7.926 18Z"/>`),
		g.Group(children),
	)
}