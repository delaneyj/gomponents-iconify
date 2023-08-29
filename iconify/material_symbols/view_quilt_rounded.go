package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewQuiltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 12Zm-10.175-.5q-.425 0-.713-.288t-.287-.712V6q0-.425.288-.713T10.825 5H20q.425 0 .713.288T21 6v4.5q0 .425-.288.713T20 11.5h-9.175Zm6.1 7.5q-.425 0-.713-.288T15.926 18v-4.5q0-.425.288-.713t.712-.287H20q.425 0 .713.288T21 13.5V18q0 .425-.288.713T20 19h-3.075Zm-6.1 0q-.425 0-.713-.288T9.825 18v-4.5q0-.425.288-.713t.712-.287h3.1q.425 0 .713.288t.287.712V18q0 .425-.288.713t-.712.287h-3.1ZM4 19q-.425 0-.713-.288T3 18V6q0-.425.288-.713T4 5h3.825q.425 0 .713.288T8.825 6v12q0 .425-.287.713T7.825 19H4Z"/>`),
		g.Group(children),
	)
}