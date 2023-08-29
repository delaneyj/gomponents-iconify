package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNoteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21q-.425 0-.713-.288T12 20v-.7q0-.2.075-.388t.225-.337l5-5l2.125 2.125l-5 5q-.15.15-.338.225T13.7 21H13Zm-9-5q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h5q.425 0 .713.288T10 15q0 .425-.288.713T9 16H4Zm16.125-1L18 12.875l.725-.725q.275-.275.7-.275t.7.275l.725.725q.275.275.275.7t-.275.7l-.725.725ZM4 12q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h9q.425 0 .713.288T14 11q0 .425-.288.713T13 12H4Zm0-4q-.425 0-.713-.288T3 7q0-.425.288-.713T4 6h9q.425 0 .713.288T14 7q0 .425-.288.713T13 8H4Z"/>`),
		g.Group(children),
	)
}