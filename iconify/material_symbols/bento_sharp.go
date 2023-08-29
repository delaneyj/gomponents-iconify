package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BentoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11V5h6v6h-6ZM2 19V5h12v14H2Zm6-5.5q.625 0 1.063-.438T9.5 12q0-.625-.438-1.063T8 10.5q-.625 0-1.063.438T6.5 12q0 .625.438 1.063T8 13.5Zm8 5.5v-6h6v6h-6Z"/>`),
		g.Group(children),
	)
}