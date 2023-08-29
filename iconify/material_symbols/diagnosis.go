package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diagnosis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19h8v-2H8v2Zm0-3h8v-2H8v2Zm4-3.7q1.65-1.5 2.825-2.663T16 7.2q0-.9-.65-1.55T13.8 5q-.525 0-1.013.213T12 5.8q-.3-.375-.788-.587T10.2 5q-.9 0-1.55.65T8 7.2q0 1.275 1.137 2.4T12 12.3Zm6 9.7H6q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22Z"/>`),
		g.Group(children),
	)
}