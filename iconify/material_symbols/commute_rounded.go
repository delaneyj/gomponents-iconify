package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommuteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.5v.5q0 .425-.288.713T11 20q-.425 0-.713-.288T10 19v-5q0-.175.025-.35t.075-.325l1.05-2.975q.2-.6.725-.975T13.05 9h5.9q.65 0 1.175.375t.725.975l1.05 2.975q.05.15.075.325T22 14v5q0 .425-.287.713T21 20q-.425 0-.713-.288T20 19v-.5h-8Zm0-6h8l-.7-2h-6.6l-.7 2Zm1 4q.425 0 .713-.288T14 15.5q0-.425-.288-.713T13 14.5q-.425 0-.713.288T12 15.5q0 .425.288.713T13 16.5Zm6 0q.425 0 .713-.288T20 15.5q0-.425-.288-.713T19 14.5q-.425 0-.713.288T18 15.5q0 .425.288.713T19 16.5ZM4.2 20q-.35 0-.475-.3t.125-.55L5 18q-1.25 0-2.125-.875T2 15V7q0-1.65 1.475-2.325T8.5 4q3.7 0 5.1.65T15 7v1h-2V7H4v6h5v7H4.2Zm.8-4q.425 0 .712-.288T6 15q0-.425-.288-.713T5 14q-.425 0-.713.288T4 15q0 .425.288.713T5 16Z"/>`),
		g.Group(children),
	)
}