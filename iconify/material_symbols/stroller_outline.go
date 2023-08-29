package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrollerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 22q-.825 0-1.413-.588T14 20q0-.825.588-1.413T16 18q.825 0 1.413.588T18 20q0 .825-.588 1.413T16 22ZM6 22q-.825 0-1.413-.588T4 20q0-.825.588-1.413T6 18q.825 0 1.413.588T8 20q0 .825-.588 1.413T6 22Zm3.6-7H15V8.675L9.6 15Zm-2.175 2q-.65 0-.913-.575t.163-1.075L16.2 4.175q.5-.575 1.113-.875T18.65 3q1.4 0 2.375.975T22 6.35V7h-2v-.65q0-.575-.388-.962T18.65 5q-.275 0-.487.1t-.388.3L17 6.275V15q0 .825-.587 1.413T15 17H7.425Zm2.05-10.35l1.375-1.6q-.2-.025-.413-.038T10 5q-.5 0-.963.075t-.937.2L9.475 6.65ZM9.6 9.6L4.725 4.725q1.2-.85 2.525-1.287T10 3q1.125 0 2.212.275T14.3 4.1L9.6 9.6Zm2.7 2.225ZM9.475 6.65Z"/>`),
		g.Group(children),
	)
}