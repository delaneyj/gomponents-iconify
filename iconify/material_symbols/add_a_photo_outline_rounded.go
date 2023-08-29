package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAPhotoOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13Zm-8 8q-.825 0-1.413-.588T1 19V7q0-.825.588-1.413T3 5h3.15L7.4 3.65q.275-.3.663-.475T8.874 3H13q.425 0 .713.288T14 4q0 .425-.288.713T13 5H8.875L7.05 7H3v12h16v-8q0-.425.288-.713T20 10q.425 0 .713.288T21 11v8q0 .825-.588 1.413T19 21H3ZM19 5h-1q-.425 0-.713-.288T17 4q0-.425.288-.713T18 3h1V2q0-.425.288-.713T20 1q.425 0 .713.288T21 2v1h1q.425 0 .713.288T23 4q0 .425-.288.713T22 5h-1v1q0 .425-.288.713T20 7q-.425 0-.713-.288T19 6V5Zm-8 12.5q1.875 0 3.188-1.313T15.5 13q0-1.875-1.313-3.188T11 8.5q-1.875 0-3.188 1.313T6.5 13q0 1.875 1.313 3.188T11 17.5Zm0-2q-1.05 0-1.775-.725T8.5 13q0-1.05.725-1.775T11 10.5q1.05 0 1.775.725T13.5 13q0 1.05-.725 1.775T11 15.5Z"/>`),
		g.Group(children),
	)
}