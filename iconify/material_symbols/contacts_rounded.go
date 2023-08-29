package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13q1.25 0 2.125-.875T15 10q0-1.25-.875-2.125T12 7q-1.25 0-2.125.875T9 10q0 1.25.875 2.125T12 13Zm-8 7q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm1.75-2h12.5q-1.125-1.4-2.725-2.2T12 15q-1.925 0-3.525.8T5.75 18ZM5 23q-.425 0-.713-.288T4 22q0-.425.288-.713T5 21h14q.425 0 .713.288T20 22q0 .425-.288.713T19 23H5ZM5 3q-.425 0-.713-.288T4 2q0-.425.288-.713T5 1h14q.425 0 .713.288T20 2q0 .425-.288.713T19 3H5Z"/>`),
		g.Group(children),
	)
}