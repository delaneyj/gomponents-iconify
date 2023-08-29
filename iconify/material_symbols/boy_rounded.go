package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.5q-.725 0-1.238-.513T10.25 5.75q0-.725.513-1.238T12 4q.725 0 1.238.513t.512 1.237q0 .725-.513 1.238T12 7.5ZM11 20q-.425 0-.713-.288T10 19v-4q-.4 0-.7-.3T9 14v-3.5q0-.825.588-1.413T11 8.5h2q.825 0 1.413.588T15 10.5V14q0 .4-.3.7t-.7.3v4q0 .425-.288.713T13 20h-2Z"/>`),
		g.Group(children),
	)
}