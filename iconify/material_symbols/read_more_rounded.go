package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadMoreRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.075 13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h6.075L6.75 8.65q-.275-.275-.288-.688t.288-.712q.275-.275.7-.275t.7.275l4.05 4.05q.3.3.3.7t-.3.7l-4.05 4.05q-.275.275-.688.288t-.712-.288q-.275-.275-.275-.7t.275-.7L9.075 13ZM14 17q-.425 0-.712-.288T13 16q0-.425.288-.713T14 15h7q.425 0 .713.288T22 16q0 .425-.288.713T21 17h-7Zm0-8q-.425 0-.712-.288T13 8q0-.425.288-.713T14 7h7q.425 0 .713.288T22 8q0 .425-.288.713T21 9h-7Zm3 4q-.425 0-.713-.288T16 12q0-.425.288-.713T17 11h4q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-4Z"/>`),
		g.Group(children),
	)
}