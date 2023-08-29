package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZonePersonUrgent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6V3q0-.825.588-1.413T4 1h3v2H4v3H2Zm5 15H4q-.825 0-1.413-.588T2 19v-3h2v3h3v2ZM20 6V3h-3V1h3q.825 0 1.413.588T22 3v3h-2Zm-6.5 2.5q-.825 0-1.413-.588T11.5 6.5q0-.825.588-1.413T13.5 4.5q.825 0 1.413.588T15.5 6.5q0 .825-.588 1.413T13.5 8.5ZM8.8 17l1-5.1l-1.8.7V16H6v-4.7l3.95-1.7q.875-.375 1.288-.487T12 9q.525 0 .975.275T13.7 10l1 1.6q.275.425.6.788t.75.662l-1.025 1.775q-.425-.275-.813-.613T13.5 13.5l-.7 3.5h-4Zm6.2 5q-.575 0-.863-.5t-.012-1l4-7q.325-.475.9-.525t.85.525l4 7q.325.5 0 1T23 22h-8Zm4-1q.2 0 .35-.15t.15-.35q0-.2-.15-.35T19 20q-.2 0-.35.15t-.15.35q0 .2.15.35T19 21Zm-.5-2h1v-4h-1v4Z"/>`),
		g.Group(children),
	)
}