package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Margin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3ZM8 9q.425 0 .713-.288T9 8q0-.425-.288-.713T8 7q-.425 0-.713.288T7 8q0 .425.288.713T8 9Zm4 0q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9Zm4 0q.425 0 .713-.288T17 8q0-.425-.288-.713T16 7q-.425 0-.713.288T15 8q0 .425.288.713T16 9Zm0 4q.425 0 .713-.288T17 12q0-.425-.288-.713T16 11q-.425 0-.713.288T15 12q0 .425.288.713T16 13Zm-4 0q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm-4 0q.425 0 .713-.288T9 12q0-.425-.288-.713T8 11q-.425 0-.713.288T7 12q0 .425.288.713T8 13Z"/>`),
		g.Group(children),
	)
}