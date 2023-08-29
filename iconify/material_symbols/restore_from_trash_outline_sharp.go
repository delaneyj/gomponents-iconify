package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestoreFromTrashOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16h2v-4.15l1.6 1.55L16 12l-4-4l-4 4l1.4 1.4l1.6-1.55V16Zm-6 5V6H4V4h5V3h6v1h5v2h-1v15H5Zm2-2h10V6H7v13ZM7 6v13V6Z"/>`),
		g.Group(children),
	)
}