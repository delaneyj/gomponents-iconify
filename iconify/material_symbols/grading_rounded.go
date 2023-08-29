package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.425 18.175L18.9 15.7q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-3.175 3.175q-.3.3-.7.3t-.7-.3L14.3 18.85q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l.725.725ZM4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h7q.425 0 .713.288T12 20q0 .425-.288.713T11 21H4Zm0-4q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h7q.425 0 .713.288T12 16q0 .425-.288.713T11 17H4Zm0-4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h16q.425 0 .713.288T21 12q0 .425-.288.713T20 13H4Zm0-4q-.425 0-.713-.288T3 8q0-.425.288-.713T4 7h16q.425 0 .713.288T21 8q0 .425-.288.713T20 9H4Zm0-4q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5H4Z"/>`),
		g.Group(children),
	)
}