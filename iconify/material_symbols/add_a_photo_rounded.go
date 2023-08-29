package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAPhotoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 7q-.425 0-.713-.288T19 6V5h-1q-.425 0-.713-.288T17 4q0-.425.288-.713T18 3h1V2q0-.425.288-.713T20 1q.425 0 .713.288T21 2v1h1q.425 0 .713.288T23 4q0 .425-.288.713T22 5h-1v1q0 .425-.288.713T20 7Zm-9 10.5q1.875 0 3.188-1.313T15.5 13q0-1.875-1.313-3.188T11 8.5q-1.875 0-3.188 1.313T6.5 13q0 1.875 1.313 3.188T11 17.5Zm0-2q-1.05 0-1.775-.725T8.5 13q0-1.05.725-1.775T11 10.5q1.05 0 1.775.725T13.5 13q0 1.05-.725 1.775T11 15.5ZM3 21q-.825 0-1.413-.588T1 19V7q0-.825.588-1.413T3 5h3.15L7.4 3.65q.275-.3.663-.475T8.874 3H14q.425 0 .713.288T15 4v1.5q0 .625.438 1.063T16.5 7h.5v.5q0 .625.438 1.063T18.5 9H20q.425 0 .713.288T21 10v9q0 .825-.588 1.413T19 21H3Z"/>`),
		g.Group(children),
	)
}