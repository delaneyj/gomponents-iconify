package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemUpdateSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23V1h14v22H5Zm2-5h10V6H7v12Zm5-2l-4-4l1.4-1.4l1.6 1.55V8h2v4.15l1.6-1.55L16 12l-4 4Z"/>`),
		g.Group(children),
	)
}