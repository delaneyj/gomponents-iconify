package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnLeftRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.8 11l.9.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.6-2.6q-.3-.3-.3-.7t.3-.7l2.6-2.6q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-.9.9H15q.825 0 1.413.588T17 11v8q0 .425-.288.713T16 20q-.425 0-.713-.288T15 19v-8H6.8Z"/>`),
		g.Group(children),
	)
}