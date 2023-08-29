package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 17v-3q0-.425.288-.713T18 13q.425 0 .713.275t.287.7V18q0 .425-.288.713T18 19H6.85l.825.825q.3.3.313.725t-.288.725q-.3.3-.713.3T6.3 21.3l-2.6-2.6q-.275-.275-.275-.7t.275-.7l2.575-2.575q.3-.3.713-.3t.712.3q.3.3.3.713t-.3.712l-.85.85H17ZM7 7v3q0 .425-.288.713T6 11q-.425 0-.713-.275t-.287-.7V6q0-.425.288-.713T6 5h11.15l-.825-.825q-.3-.3-.313-.725t.288-.725q.3-.3.713-.3t.687.275l2.6 2.6q.275.275.275.7t-.275.7l-2.575 2.575q-.3.3-.713.3t-.712-.3q-.3-.3-.3-.713t.3-.712l.85-.85H7Z"/>`),
		g.Group(children),
	)
}