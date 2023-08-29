package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueueMusicRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20q-1.25 0-2.125-.875T13 17q0-1.25.875-2.125T16 14q.275 0 .525.038T17 14.2V7q0-.425.288-.713T18 6h3q.425 0 .713.288T22 7q0 .425-.288.713T21 8h-2v9q0 1.25-.875 2.125T16 20ZM4 16q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h6q.425 0 .713.288T11 15q0 .425-.288.713T10 16H4Zm0-4q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h10q.425 0 .713.288T15 11q0 .425-.288.713T14 12H4Zm0-4q-.425 0-.713-.288T3 7q0-.425.288-.713T4 6h10q.425 0 .713.288T15 7q0 .425-.288.713T14 8H4Z"/>`),
		g.Group(children),
	)
}