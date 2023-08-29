package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsFlatHead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2h8v2H8Zm0-3l-1-7l2-8h6l2 8l-1 7H8Zm1.3-8h5.4l-1.25-5h-2.9L9.3 10Z"/>`),
		g.Group(children),
	)
}