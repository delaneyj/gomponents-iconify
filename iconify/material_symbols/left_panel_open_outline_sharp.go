package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftPanelOpenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 8v8l4-4l-4-4ZM5 19h3V5H5v14Zm5 0h9V5h-9v14Zm-2 0H5h3Zm-5 2V3h18v18H3Z"/>`),
		g.Group(children),
	)
}