package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscoverTuneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 9q-.425 0-.713-.288T13 8q0-.425.288-.713T14 7h2V4q0-.425.288-.713T17 3q.425 0 .713.288T18 4v3h2q.425 0 .713.288T21 8q0 .425-.288.713T20 9h-6Zm3 12q-.425 0-.713-.288T16 20v-8q0-.425.288-.713T17 11q.425 0 .713.288T18 12v8q0 .425-.288.713T17 21ZM7 21q-.425 0-.713-.288T6 20v-3H4q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h6q.425 0 .713.288T11 16q0 .425-.288.713T10 17H8v3q0 .425-.288.713T7 21Zm0-8q-.425 0-.713-.288T6 12V4q0-.425.288-.713T7 3q.425 0 .713.288T8 4v8q0 .425-.288.713T7 13Z"/>`),
		g.Group(children),
	)
}