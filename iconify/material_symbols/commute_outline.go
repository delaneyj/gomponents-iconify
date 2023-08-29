package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommuteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 19.35v-5.7l1.4-4q.125-.275.363-.463T12.35 9h7.3q.35 0 .6.188t.35.462l1.4 4v5.7q0 .275-.187.463T21.35 20h-.7q-.275 0-.462-.188T20 19.35v-.85h-8v.85q0 .275-.188.463T11.35 20h-.7q-.275 0-.462-.188T10 19.35Zm2-6.85h8l-.7-2h-6.6l-.7 2Zm-.5 1.5v3v-3Zm1.5 2.5q.425 0 .713-.288T14 15.5q0-.425-.288-.713T13 14.5q-.425 0-.713.288T12 15.5q0 .425.288.713T13 16.5Zm6 0q.425 0 .713-.288T20 15.5q0-.425-.288-.713T19 14.5q-.425 0-.713.288T18 15.5q0 .425.288.713T19 16.5ZM4 20v-1l1-1q-1.25 0-2.125-.875T2 15V7q0-1.65 1.475-2.325T8.5 4q3.7 0 5.1.65T15 7v1h-2V7H4v6h5v7H4Zm1-4q.425 0 .713-.288T6 15q0-.425-.288-.713T5 14q-.425 0-.713.288T4 15q0 .425.288.713T5 16Zm6.5 1h9v-3h-9v3Z"/>`),
		g.Group(children),
	)
}