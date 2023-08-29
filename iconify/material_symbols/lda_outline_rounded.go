package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LdaOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 18.575l-4.775-2q-.575-.225-.9-.738T5 14.725V13q0-.825.588-1.412T7 11h4V8h-1q-.825 0-1.413-.588T8 6V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v2q0 .825-.588 1.413T14 8h-1v3h4q.825 0 1.413.588T19 13v1.725q0 .6-.325 1.113t-.9.737l-4.775 2V21q0 .425-.287.713T12 22q-.425 0-.713-.288T11 21v-2.425ZM10 6h4V4h-4v2Zm1 10.4V13H7v1.725l4 1.675Zm2 0l4-1.675V13h-4v3.4ZM10 6V4v2Z"/>`),
		g.Group(children),
	)
}