package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpToElementRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11h1q.425 0 .713.288T15 12q0 .425-.288.713T14 13h-2q-.425 0-.713-.288T11 12v-2q0-.425.288-.713T12 9q.425 0 .713.288T13 10v1Zm7 0v-1q0-.425.288-.713T21 9q.425 0 .713.288T22 10v2q0 .425-.288.713T21 13h-2q-.425 0-.713-.288T18 12q0-.425.288-.713T19 11h1Zm-7-7v1q0 .425-.288.713T12 6q-.425 0-.713-.288T11 5V3q0-.425.288-.713T12 2h2q.425 0 .713.288T15 3q0 .425-.288.713T14 4h-1Zm7 0h-1q-.425 0-.713-.288T18 3q0-.425.288-.713T19 2h2q.425 0 .713.288T22 3v2q0 .425-.288.713T21 6q-.425 0-.713-.288T20 5V4ZM9 16.4l-4.9 4.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7L7.6 15H4q-.425 0-.713-.288T3 14q0-.425.288-.713T4 13h6q.425 0 .713.288T11 14v6q0 .425-.288.713T10 21q-.425 0-.713-.288T9 20v-3.6Z"/>`),
		g.Group(children),
	)
}