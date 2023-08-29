package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteAddOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15v2q0 .425.288.713T12 18q.425 0 .713-.288T13 17v-2h2q.425 0 .713-.288T16 14q0-.425-.288-.713T15 13h-2v-2q0-.425-.288-.713T12 10q-.425 0-.713.288T11 11v2H9q-.425 0-.713.288T8 14q0 .425.288.713T9 15h2Zm-5 7q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V20q0 .825-.588 1.413T18 22H6Zm7-14V4H6v16h12V9h-4q-.425 0-.713-.288T13 8ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}