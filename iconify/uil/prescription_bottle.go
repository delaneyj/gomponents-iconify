package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrescriptionBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-2 14h-6v-4h6Zm0-6h-7a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h7v1a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8h10ZM5 6V4h14v2Z"/>`),
		g.Group(children),
	)
}