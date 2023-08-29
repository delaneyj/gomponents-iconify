package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterNinePlusOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 14h4V6H9v5h3v1h-2v2Zm2-5h-1V8h1v1Zm-6 9V2h16v16H6Zm2-2h12V4H8v12Zm-6 6V6h2v14h14v2H2Zm6-6V4v12Zm8.5-3h2v-2H20V9h-1.5V7h-2v2h-2v2h2v2Z"/>`),
		g.Group(children),
	)
}