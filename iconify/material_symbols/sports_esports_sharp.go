package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsEsportsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1.625 19l1.95-14h16.85l1.95 14H18.4l-3-3H8.6l-3 3H1.625ZM17 13q.425 0 .713-.288T18 12q0-.425-.288-.713T17 11q-.425 0-.713.288T16 12q0 .425.288.713T17 13Zm-2-3q.425 0 .713-.288T16 9q0-.425-.288-.713T15 8q-.425 0-.713.288T14 9q0 .425.288.713T15 10Zm-7.25 3h1.5v-1.75H11v-1.5H9.25V8h-1.5v1.75H6v1.5h1.75V13Z"/>`),
		g.Group(children),
	)
}