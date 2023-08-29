package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridGoldenratioRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15H3q-.425 0-.713-.288T2 14q0-.425.288-.713T3 13h6v-2H3q-.425 0-.713-.288T2 10q0-.425.288-.713T3 9h6V3q0-.425.288-.713T10 2q.425 0 .713.288T11 3v6h2V3q0-.425.288-.713T14 2q.425 0 .713.288T15 3v6h6q.425 0 .713.288T22 10q0 .425-.288.713T21 11h-6v2h6q.425 0 .713.288T22 14q0 .425-.288.713T21 15h-6v6q0 .425-.288.713T14 22q-.425 0-.713-.288T13 21v-6h-2v6q0 .425-.288.713T10 22q-.425 0-.713-.288T9 21v-6Zm2-2h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}