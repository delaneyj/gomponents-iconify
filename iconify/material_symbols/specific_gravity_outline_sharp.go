package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpecificGravityOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.2 22L3 2h18l-2.2 20H5.2Zm.475-14H9.35q.525-.475 1.2-.738T12 7q.775 0 1.45.263t1.2.737h3.675l.425-4H5.225l.45 4ZM12 13q.825 0 1.413-.588T14 11q0-.825-.588-1.413T12 9q-.825 0-1.413.588T10 11q0 .825.588 1.413T12 13Zm-6.1-3L7 20h10l1.1-10h-2.225q.05.25.088.488T16 11q0 1.65-1.175 2.825T12 15q-1.65 0-2.825-1.175T8 11q0-.275.038-.513T8.124 10H5.9ZM7 20h1.125H8h8h-.125H17H7Z"/>`),
		g.Group(children),
	)
}