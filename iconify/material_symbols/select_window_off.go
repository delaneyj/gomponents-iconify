package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectWindowOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L10.175 13H4v7h12v-4.025l2 2V20q0 .825-.588 1.413T16 22H4q-.825 0-1.413-.588T2 20v-9q0-.825.588-1.413T4 9h2v-.175L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM18 15.125L15.875 13l-4-4H16q.825 0 1.413.588T18 11v2h2V6H8.875L6.15 3.275q.2-.575.713-.925T8 2h12q.825 0 1.413.588T22 4v9q0 .825-.588 1.413T20 15h-2v.125Z"/>`),
		g.Group(children),
	)
}