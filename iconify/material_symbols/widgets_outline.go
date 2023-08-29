package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidgetsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.65 13L11 7.35l5.65-5.65l5.65 5.65L16.65 13ZM3 11V3h8v8H3Zm10 10v-8h8v8h-8ZM3 21v-8h8v8H3ZM5 9h4V5H5v4Zm11.675 1.2L19.5 7.375L16.675 4.55L13.85 7.375l2.825 2.825ZM15 19h4v-4h-4v4ZM5 19h4v-4H5v4ZM9 9Zm4.85-1.625ZM9 15Zm6 0Z"/>`),
		g.Group(children),
	)
}