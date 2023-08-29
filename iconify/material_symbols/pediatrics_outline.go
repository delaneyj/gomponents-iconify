package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PediatricsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7q-.425 0-.713-.288T7 6q0-.425.288-.713T8 5h3V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v2h3q.425 0 .713.288T17 6q0 .425-.288.713T16 7H8Zm1 15q-.825 0-1.413-.588T7 20v-9q0-1.25.875-2.125T10 8h4q1.25 0 2.125.875T17 11v9q0 .825-.588 1.413T15 22H9Zm0-2h6v-9q0-.425-.288-.713T14 10h-4q-.425 0-.713.288T9 11v1h2q.425 0 .713.288T12 13q0 .425-.288.713T11 14H9v2h2q.425 0 .713.288T12 17q0 .425-.288.713T11 18H9v2Zm0 0V10v10Z"/>`),
		g.Group(children),
	)
}