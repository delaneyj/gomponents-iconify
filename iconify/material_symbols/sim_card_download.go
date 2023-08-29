package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17l4-4l-1.4-1.4l-1.6 1.55V9h-2v4.15L9.4 11.6L8 13l4 4Zm-6 5q-.825 0-1.413-.588T4 20V8l6-6h8q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Z"/>`),
		g.Group(children),
	)
}