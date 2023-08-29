package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiTransportationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22q-.425 0-.713-.288T10 21v-5q0-.175.025-.35t.075-.325l1.05-2.975q.2-.6.725-.975T13.05 11h5.9q.65 0 1.175.375t.725.975l1.05 2.975q.05.15.075.325T22 16v5q0 .425-.287.713T21 22q-.425 0-.713-.288T20 21v-.5h-8v.5q0 .425-.288.713T11 22Zm1-7.5h8l-.7-2h-6.6l-.7 2Zm1 4q.425 0 .713-.288T14 17.5q0-.425-.288-.713T13 16.5q-.425 0-.713.288T12 17.5q0 .425.288.713T13 18.5Zm6 0q.425 0 .713-.288T20 17.5q0-.425-.288-.713T19 16.5q-.425 0-.713.288T18 17.5q0 .425.288.713T19 18.5ZM6 14v-2h2v2H6Zm5-6V6h2v2h-2ZM6 18v-2h2v2H6Zm0 4v-2h2v2H6Zm-4 0V10q0-.825.588-1.413T4 8h3V4q0-.825.588-1.413T9 2h6q.825 0 1.413.588T17 4v5h-2V4H9v6H4v12H2Z"/>`),
		g.Group(children),
	)
}