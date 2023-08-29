package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermCameraMicOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21H2V5h5.15L9 3h6l1.85 2H22v16h-7v-2h5V7h-4.05l-1.825-2h-4.25L8.05 7H4v12h5v2Zm2 0h2v-3.1q2.15-.35 3.575-2.012T18 12h-2q0 1.65-1.175 2.825T12 16q-1.65 0-2.825-1.175T8 12H6q0 2.225 1.425 3.888T11 17.9V21Zm1-7q.825 0 1.413-.588T14 12V8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8v4q0 .825.588 1.413T12 14Zm-8 5h16H4Z"/>`),
		g.Group(children),
	)
}