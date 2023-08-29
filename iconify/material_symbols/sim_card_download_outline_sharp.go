package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardDownloadOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17l4-4l-1.4-1.4l-1.6 1.55V9h-2v4.15L9.4 11.6L8 13l4 4Zm-8 5V8l6-6h10v20H4Zm2-2h12V4h-7.15L6 8.85V20Zm0 0h12H6Z"/>`),
		g.Group(children),
	)
}