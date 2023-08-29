package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubdirectoryArrowLeftRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.825 16l2.875 2.875q.3.3.313.7t-.288.7q-.3.3-.712.3t-.713-.3L4.7 15.7q-.15-.15-.213-.325T4.425 15q0-.2.063-.375T4.7 14.3l4.625-4.625q.3-.3.7-.287t.7.312q.275.3.287.7t-.287.7l-2.9 2.9H17V5q0-.425.288-.713T18 4q.425 0 .713.288T19 5v9q0 .825-.588 1.413T17 16H7.825Z"/>`),
		g.Group(children),
	)
}