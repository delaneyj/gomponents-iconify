package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsPhillipsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-2h10v2H7Zm0-3V8l3.75-5h2.5L17 8v10H7Zm2-2h6v-2.175l-3-3l-3 3V16Zm0-5l2-2V6L9 8.675V11Zm6 0V8.675L13 6v3l2 2Zm0 5H9h6Z"/>`),
		g.Group(children),
	)
}