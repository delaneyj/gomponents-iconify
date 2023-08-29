package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSidebarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2ZM17.5 8.675H20V6h-2.5v2.675Zm0 4.65H20v-2.65h-2.5v2.65ZM4 18h11.5V6H4v12Zm13.5 0H20v-2.675h-2.5V18Z"/>`),
		g.Group(children),
	)
}