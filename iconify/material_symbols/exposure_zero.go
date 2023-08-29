package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposureZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q-2.5 0-4-1.988T6.5 12q0-3.025 1.5-5.013T12 5q2.5 0 4 1.988T17.5 12q0 3.025-1.5 5.013T12 19Zm0-2.05q1.65 0 2.475-1.5T15.3 12q0-1.95-.825-3.45T12 7.05q-1.65 0-2.475 1.5T8.7 12q0 1.95.825 3.45T12 16.95Z"/>`),
		g.Group(children),
	)
}