package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardTabRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18q-.425 0-.713-.288T20 17V7q0-.425.288-.713T21 6q.425 0 .713.288T22 7v10q0 .425-.288.713T21 18Zm-6.825-5H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h11.175L11.3 8.1q-.275-.275-.288-.688T11.3 6.7q.275-.275.7-.275t.7.275l4.6 4.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-4.6 4.6q-.275.275-.687.275T11.3 17.3q-.3-.3-.3-.713t.3-.712L14.175 13Z"/>`),
		g.Group(children),
	)
}