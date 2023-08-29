package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SheetsRtlSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 13V1h16v12H4Zm2-7h5V3H6v3Zm7 0h5V3h-5v3Zm-7 5h5V8H6v3Zm7 0h5V8h-5v3ZM7 22l-4-4l4-4l1.4 1.4L6.825 17H20v2H6.825L8.4 20.6L7 22Z"/>`),
		g.Group(children),
	)
}