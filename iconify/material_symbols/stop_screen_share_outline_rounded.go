package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopScreenShareOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.425 11.575L11.85 9H13V7.6q0-.175.15-.238t.275.063L15.65 9.65q.15.15.15.35t-.15.35l-1.225 1.225ZM20.7 17.85L18.85 16H20V5H7.85l-2-2H20q.825 0 1.413.588T22 5v11q0 .65-.363 1.137t-.937.713ZM18.2 21H2q-.425 0-.713-.288T1 20q0-.425.288-.713T2 19h14.175l-1-1H4q-.825 0-1.413-.588T2 16V4.85l-.625-.65Q1.1 3.9 1.1 3.5t.3-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L18.2 21ZM10 12.85V13q0 .425-.288.713T9 14q-.425 0-.713-.288T8 13v-1q0-.275.025-.525t.15-.475L4 6.825V16h9.15L10 12.85Zm3.35-2.35Zm-4.775.9Z"/>`),
		g.Group(children),
	)
}