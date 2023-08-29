package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentsDisabledOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V4.825L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3L15.175 18H4Zm18 1.125L18.875 16H20V4H6.875l-2-2H20q.825 0 1.413.588T22 4v15.125ZM4 16h9.175l-2-2H7q-.425 0-.713-.288T6 13q0-.425.288-.713T7 12h2.175l-1-1H7q-.425 0-.712-.288T6 10q0-.425.288-.713T7 9h.625v1.45L4 6.825V16Zm12.875-2l-2-2H17q.425 0 .713.288T18 13q0 .425-.288.713T17 14h-.125Zm-3-3l-2-2H17q.425 0 .713.288T18 10q0 .425-.288.713T17 11h-3.125Zm-3-3l-2-2H17q.425 0 .713.288T18 7q0 .425-.288.713T17 8h-6.125ZM8.6 11.4Zm4.275-1.4Z"/>`),
		g.Group(children),
	)
}