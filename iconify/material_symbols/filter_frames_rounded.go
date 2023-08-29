package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFramesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V6q0-.825.588-1.413T4 4h4L11.3.7q.3-.3.7-.3t.7.3L16 4h4q.825 0 1.413.588T22 6v14q0 .825-.588 1.413T20 22H4Zm0-2h16V6H4v14Zm2-3V9q0-.425.288-.713T7 8h10q.425 0 .713.288T18 9v8q0 .425-.288.713T17 18H7q-.425 0-.713-.288T6 17Z"/>`),
		g.Group(children),
	)
}