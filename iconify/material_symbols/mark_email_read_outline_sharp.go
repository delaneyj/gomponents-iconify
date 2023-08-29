package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkEmailReadOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.95 22l-4.25-4.25l1.4-1.4l2.85 2.85l5.65-5.65l1.4 1.4L15.95 22ZM12 11l8-5H4l8 5Zm0 2L4 8v10h5.15l2 2H2V4h20v6.35l-2 2V8l-8 5Zm0 0Zm0-2Zm0 2Z"/>`),
		g.Group(children),
	)
}