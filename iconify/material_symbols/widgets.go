package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Widgets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.65 13L11 7.35l5.65-5.65l5.65 5.65L16.65 13ZM3 11V3h8v8H3Zm10 10v-8h8v8h-8ZM3 21v-8h8v8H3Z"/>`),
		g.Group(children),
	)
}