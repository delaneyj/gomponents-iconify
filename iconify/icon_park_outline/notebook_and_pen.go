package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookAndPen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 6v36h26V6H4Zm8 36V6m32 0h-8v32l4 4l4-4V6Zm-8 6h8M30 6H4m26 36H4M36 6v16m8-16v16"/>`),
		g.Group(children),
	)
}