package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPantsSweat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 19l9 19h9L38 4H10L6 38h9l9-19Zm10 19l1 6h6v-6h-7Zm-21 6H7v-6h7l-1 6ZM24 4l4 7.5M24 4l-4 7.5"/>`),
		g.Group(children),
	)
}