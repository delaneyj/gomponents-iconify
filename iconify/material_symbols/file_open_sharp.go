package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOpenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.95 22.375L19 19.425v2.225h-2V16h5.65v2H20.4l2.95 2.95l-1.4 1.425ZM13 9h5l-5-5v5ZM4 22V2h10l6 6v6h-5v8H4Z"/>`),
		g.Group(children),
	)
}