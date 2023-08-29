package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SellOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.825 22.825L2 12V2h10l10.8 10.85l-9.975 9.975Zm0-2.825l7.15-7.15L11.15 4H4v7.15L12.825 20ZM6.5 8q.625 0 1.063-.438T8 6.5q0-.625-.438-1.063T6.5 5q-.625 0-1.063.438T5 6.5q0 .625.438 1.063T6.5 8ZM4 4Z"/>`),
		g.Group(children),
	)
}