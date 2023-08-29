package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.825 12.025L8.7 15.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-4.6-4.6q-.15-.15-.213-.325T2.426 12q0-.2.063-.375T2.7 11.3l2.875-2.875l-3.5-3.5q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L7 9.85l-2.175 2.175Zm13.6 3.55L17 14.15l2.175-2.175L15.3 8.1q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l4.6 4.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-2.875 2.875Z"/>`),
		g.Group(children),
	)
}