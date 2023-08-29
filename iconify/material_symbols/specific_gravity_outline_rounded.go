package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpecificGravityOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-.775 0-1.35-.5T5 20.225L3.125 3.1q-.05-.45.25-.775T4.125 2h15.75q.45 0 .75.325t.25.775L19 20.225q-.075.775-.65 1.275T17 22H7ZM5.675 8H9.35q.525-.475 1.2-.738T12 7q.775 0 1.45.263t1.2.737h3.675l.425-4H5.225l.45 4ZM12 13q.825 0 1.413-.588T14 11q0-.825-.588-1.413T12 9q-.825 0-1.413.588T10 11q0 .825.588 1.413T12 13Zm-6.1-3L7 20h10l1.1-10h-2.225q.05.25.088.488T16 11q0 1.65-1.175 2.825T12 15q-1.65 0-2.825-1.175T8 11q0-.275.038-.513T8.124 10H5.9ZM7 20h1.125H8h8h-.125H17H7Z"/>`),
		g.Group(children),
	)
}