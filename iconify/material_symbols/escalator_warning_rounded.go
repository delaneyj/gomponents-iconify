package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorWarningRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 6q-.825 0-1.413-.588T4.5 4q0-.825.588-1.413T6.5 2q.825 0 1.413.588T8.5 4q0 .825-.588 1.413T6.5 6ZM17 11q-.625 0-1.063-.438T15.5 9.5q0-.625.438-1.063T17 8q.625 0 1.063.438T18.5 9.5q0 .625-.438 1.063T17 11ZM5.5 22q-.425 0-.713-.288T4.5 21v-6H4q-.425 0-.713-.288T3 14V9q0-.825.588-1.413T5 7h3q.55 0 1 .263T9.725 8l3.575 6.175l1.025-1.525q.2-.3.538-.475t.712-.175H18.5q.625 0 1.063.438T20 13.5V16q0 .425-.288.713T19 17v4q0 .425-.288.713T18 22h-2q-.425 0-.713-.288T15 21v-6.1l-.475.675q-.15.2-.363.313T13.7 16h-1.1q-.275 0-.512-.138t-.363-.362L9.5 11.6V21q0 .425-.287.713T8.5 22h-3Z"/>`),
		g.Group(children),
	)
}