package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11q.425 0 .713-.288T9 10q0-.425-.288-.713T8 9q-.425 0-.713.288T7 10q0 .425.288.713T8 11Zm4 0q.425 0 .713-.288T13 10q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10q0 .425.288.713T12 11Zm4 0q.425 0 .713-.288T17 10q0-.425-.288-.713T16 9q-.425 0-.713.288T15 10q0 .425.288.713T16 11ZM2 22V2h20v16H6l-4 4Z"/>`),
		g.Group(children),
	)
}