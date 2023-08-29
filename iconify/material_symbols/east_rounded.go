package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EastRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.175 13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h15.175L14.3 7.1q-.275-.275-.288-.687T14.3 5.7q.275-.275.7-.275t.7.275l5.6 5.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-5.6 5.6q-.275.275-.687.275T14.3 18.3q-.3-.3-.3-.713t.3-.712L18.175 13Z"/>`),
		g.Group(children),
	)
}