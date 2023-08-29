package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoyaltySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.825 22.825L2 12V2h10l10.8 10.85l-9.975 9.975ZM6.5 8q.625 0 1.063-.438T8 6.5q0-.625-.438-1.063T6.5 5q-.625 0-1.063.438T5 6.5q0 .625.438 1.063T6.5 8Zm6.5 9.5l3.5-3.5q.275-.275.438-.65t.162-.8q0-.85-.6-1.45t-1.45-.6q-.475 0-.938.275T13 11.7q-.75-.7-1.175-.95t-.875-.25q-.85 0-1.45.6t-.6 1.45q0 .425.163.8T9.5 14l3.5 3.5Z"/>`),
		g.Group(children),
	)
}