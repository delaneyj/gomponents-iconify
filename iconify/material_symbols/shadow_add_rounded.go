package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShadowAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11h-2q-.425 0-.713-.288T10 10q0-.425.288-.713T11 9h2V7q0-.425.288-.713T14 6q.425 0 .713.288T15 7v2h2q.425 0 .713.288T18 10q0 .425-.288.713T17 11h-2v2q0 .425-.288.713T14 14q-.425 0-.713-.288T13 13v-2ZM4 22q-.825 0-1.413-.588T2 20V8q0-.825.588-1.413T4 6h2V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18h-2v2q0 .825-.588 1.413T16 22H4Zm4-6h12V4H8v12Z"/>`),
		g.Group(children),
	)
}