package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrollerRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.775 3.6q.775-.3 1.575-.45T10 3q1.125 0 2.213.275T14.3 4.1L9.6 9.6L6.375 6.375q-.3-.3-.45-.65t-.15-.7q0-.45.25-.838t.75-.587ZM16 22q-.825 0-1.413-.587T14 20q0-.825.588-1.413T16 18q.825 0 1.413.588T18 20q0 .825-.588 1.413T16 22ZM6 22q-.825 0-1.413-.587T4 20q0-.825.588-1.413T6 18q.825 0 1.413.588T8 20q0 .825-.588 1.413T6 22Zm1.425-5q-.65 0-.913-.575t.163-1.075L16.2 4.175q.5-.575 1.113-.875T18.65 3q1.4 0 2.375.975T22 6.35q0 .425-.288.713T21 7.35q-.425 0-.713-.287T20 6.35q0-.575-.388-.962T18.65 5q-.275 0-.487.1t-.388.3L17 6.275V15q0 .825-.587 1.413T15 17H7.425Z"/>`),
		g.Group(children),
	)
}