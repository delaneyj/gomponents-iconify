package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpecificGravity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-.775 0-1.35-.5T5 20.225L3 2h18l-2 18.225q-.075.775-.65 1.275T17 22H7ZM5.675 8H9.35q.525-.475 1.2-.738T12 7q.775 0 1.45.263t1.2.737h3.675l.425-4H5.225l.45 4ZM12 13q.825 0 1.413-.588T14 11q0-.825-.588-1.413T12 9q-.825 0-1.413.588T10 11q0 .825.588 1.413T12 13Z"/>`),
		g.Group(children),
	)
}