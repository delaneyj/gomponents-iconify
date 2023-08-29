package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.175 11L15.3 7.125q-.3-.3-.3-.713t.3-.712q.3-.275.713-.275t.687.275l4.6 4.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-4.6 4.6q-.3.3-.7.288t-.7-.288q-.3-.3-.313-.7t.288-.7l3.9-3.9Zm-6 1H7q-1.25 0-2.125.875T4 15v3q0 .425-.288.713T3 19q-.425 0-.713-.288T2 18v-3q0-2.075 1.463-3.538T7 10h6.175L10.3 7.125q-.3-.3-.3-.713t.3-.712q.3-.275.713-.275t.687.275l4.6 4.6q.15.15.213.325t.062.375q0 .2-.063.375t-.212.325l-4.6 4.6q-.3.3-.7.288t-.7-.288q-.3-.3-.313-.7t.288-.7l2.9-2.9Z"/>`),
		g.Group(children),
	)
}