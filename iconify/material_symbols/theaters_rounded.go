package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheatersRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19v1q0 .425-.288.713T5 21q-.425 0-.713-.288T4 20V4q0-.425.288-.713T5 3q.425 0 .713.288T6 4v1h2V4q0-.425.288-.713T9 3h6q.425 0 .713.288T16 4v1h2V4q0-.425.288-.713T19 3q.425 0 .713.288T20 4v16q0 .425-.288.713T19 21q-.425 0-.713-.288T18 20v-1h-2v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20v-1H6Zm0-2h2v-2H6v2Zm0-4h2v-2H6v2Zm0-4h2V7H6v2Zm10 8h2v-2h-2v2Zm0-4h2v-2h-2v2Zm0-4h2V7h-2v2Z"/>`),
		g.Group(children),
	)
}