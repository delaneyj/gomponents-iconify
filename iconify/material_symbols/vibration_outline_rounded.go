package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibrationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 9q.425 0 .713.288T2 10v4q0 .425-.288.713T1 15q-.425 0-.713-.288T0 14v-4q0-.425.288-.713T1 9Zm3-2q.425 0 .713.288T5 8v8q0 .425-.288.713T4 17q-.425 0-.713-.288T3 16V8q0-.425.288-.713T4 7Zm19 2q.425 0 .713.288T24 10v4q0 .425-.288.713T23 15q-.425 0-.713-.288T22 14v-4q0-.425.288-.713T23 9Zm-3-2q.425 0 .713.288T21 8v8q0 .425-.288.713T20 17q-.425 0-.713-.288T19 16V8q0-.425.288-.713T20 7ZM8 21q-.825 0-1.413-.588T6 19V5q0-.825.588-1.413T8 3h8q.825 0 1.413.588T18 5v14q0 .825-.588 1.413T16 21H8Zm0-2h8V5H8v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}