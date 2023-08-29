package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyFullscreenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm10-8v1q0 .425.288.713T15 14q.425 0 .713-.288T16 13v-1h1q.425 0 .713-.288T18 11q0-.425-.288-.713T17 10h-1V9q0-.425-.288-.713T15 8q-.425 0-.713.288T14 9v1h-1q-.425 0-.713.288T12 11q0 .425.288.713T13 12h1Z"/>`),
		g.Group(children),
	)
}