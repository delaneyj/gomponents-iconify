package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterListOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.075 21.9l-17-16.975q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3Zm-3.25-8.9l-2-2H17q.425 0 .713.288T18 12q0 .425-.288.713T17 13h-1.175Zm-5-5l-2-2H20q.425 0 .713.288T21 7q0 .425-.288.713T20 8h-9.175ZM11 18q-.425 0-.712-.288T10 17q0-.425.288-.713T11 16h2q.425 0 .713.288T14 17q0 .425-.288.713T13 18h-2Zm-.85-5H7q-.425 0-.713-.288T6 12q0-.425.288-.713T7 11h3.15v2Zm-5-5H4q-.425 0-.713-.288T3 7q0-.425.288-.713T4 6h1.15v2Z"/>`),
		g.Group(children),
	)
}