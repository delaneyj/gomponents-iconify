package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOpenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h10l6 6v6h-2V9h-5V4H6v16h9v2H4Zm17.95.375L19 19.425v2.225h-2V16h5.65v2H20.4l2.95 2.95l-1.4 1.425ZM6 20V4v16Z"/>`),
		g.Group(children),
	)
}