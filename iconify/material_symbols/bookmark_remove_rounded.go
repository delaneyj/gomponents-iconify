package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkRemoveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 7q-.425 0-.713-.288T15 6q0-.425.288-.713T16 5h4q.425 0 .713.288T21 6q0 .425-.288.713T20 7h-4Zm-4 11l-4.2 1.8q-1 .425-1.9-.163T5 17.976V5q0-.825.588-1.413T7 3h7q-.5.75-.75 1.438T13 6q0 1.8 1.137 3.175T17 10.9q.575.075 1 .075t1-.075v7.075q0 1.075-.9 1.663t-1.9.162L12 18Z"/>`),
		g.Group(children),
	)
}