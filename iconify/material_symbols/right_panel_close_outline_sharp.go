package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightPanelCloseOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.5 16l4-4l-4-4v8Zm8.5 3h3V5h-3v14ZM5 19h9V5H5v14Zm11 0h3h-3ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}