package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 9.75l1.325-.65q.325-.15.675-.15t.675.15L14 9.75V5h-4v4.75ZM8 17q-.425 0-.713-.288T7 16q0-.425.288-.713T8 15h3q.425 0 .713.288T12 16q0 .425-.288.713T11 17H8Zm-3 4q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5ZM5 5v14V5Zm0 14h14V5h-3v6.375q0 .575-.475.863t-.975.037L12 11l-2.55 1.275q-.5.25-.975-.038T8 11.376V5H5v14Z"/>`),
		g.Group(children),
	)
}