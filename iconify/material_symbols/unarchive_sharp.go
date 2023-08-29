package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnarchiveSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 10l-4 4l1.4 1.4l1.6-1.6V18h2v-4.2l1.6 1.6L16 14l-4-4ZM3 21V5.8L5.3 3h13.4L21 5.8V21H3ZM5.4 6h13.2l-.85-1H6.25L5.4 6Z"/>`),
		g.Group(children),
	)
}