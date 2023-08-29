package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommuteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.5V20h-2v-6.35L11.625 9h8.75L22 13.65V20h-2v-1.5h-8Zm0-6h8l-.7-2h-6.6l-.7 2Zm-.5 1.5v3v-3Zm1.5 2.5q.425 0 .713-.288T14 15.5q0-.425-.288-.713T13 14.5q-.425 0-.713.288T12 15.5q0 .425.288.713T13 16.5Zm6 0q.425 0 .713-.288T20 15.5q0-.425-.288-.713T19 14.5q-.425 0-.713.288T18 15.5q0 .425.288.713T19 16.5ZM4 20v-1l1-1q-1.25 0-2.125-.875T2 15V7q0-1.65 1.475-2.325T8.5 4q3.7 0 5.1.65T15 7v1h-2V7H4v6h5v7H4Zm1-4q.425 0 .713-.288T6 15q0-.425-.288-.713T5 14q-.425 0-.713.288T4 15q0 .425.288.713T5 16Zm6.5 1h9v-3h-9v3Z"/>`),
		g.Group(children),
	)
}