package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentsDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.175 18H4q-.825 0-1.413-.588T2 16V4.825L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3L15.175 18ZM22 19.125L16.875 14H17q.425 0 .713-.288T18 13q0-.425-.288-.713T17 12h-2.125l-1-1H17q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9h-5.125l-1-1H17q.425 0 .713-.288T18 7q0-.425-.288-.713T17 6H8.875l-4-4H20q.825 0 1.413.588T22 4v15.125ZM11.175 14l-2-2H7q-.425 0-.713.288T6 13q0 .425.288.713T7 14h4.175Zm-3-3l-1.8-1.8q-.175.125-.275.338T6 10q0 .425.288.713T7 11h1.175Z"/>`),
		g.Group(children),
	)
}