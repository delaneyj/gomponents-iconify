package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21v-8h8v8h-8Zm0-10V3h8v8h-8ZM3 11V3h8v8H3Zm0 10v-8h8v8H3Z"/>`),
		g.Group(children),
	)
}