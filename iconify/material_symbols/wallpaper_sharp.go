package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallpaperSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-8h2v6h6v2H3Zm10 0v-2h6v-6h2v8h-8Zm-7-4l3-4l2.25 3l3-4L18 17H6Zm-3-6V3h8v2H5v6H3Zm16 0V5h-6V3h8v8h-2Zm-3.5-1q-.65 0-1.075-.425T14 8.5q0-.65.425-1.075T15.5 7q.65 0 1.075.425T17 8.5q0 .65-.425 1.075T15.5 10Z"/>`),
		g.Group(children),
	)
}