package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm4.45-8.725L12 11l2.55 1.275q.5.25.975-.038t.475-.862V5H8v6.375q0 .575.475.863t.975.037ZM8 17h3q.425 0 .713-.288T12 16q0-.425-.288-.713T11 15H8q-.425 0-.713.288T7 16q0 .425.288.713T8 17Z"/>`),
		g.Group(children),
	)
}