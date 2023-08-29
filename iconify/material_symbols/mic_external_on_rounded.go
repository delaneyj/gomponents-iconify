package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicExternalOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.8 7q-.375-.425-.587-.925T4 5q0-1.25.875-2.125T7 2q1.25 0 2.125.875T10 5q0 .575-.213 1.075T9.2 7H4.8ZM10 22q-1.65 0-2.825-1.175T6 18h-.55q-.2 0-.338-.125t-.162-.325L4.1 9.1q-.05-.45.25-.775T5.1 8h3.8q.45 0 .75.325t.25.775l-.85 8.45q-.025.2-.163.325T8.55 18H8q0 .825.588 1.413T10 20q.825 0 1.413-.588T12 18V6q0-1.65 1.175-2.825T16 2q1.65 0 2.825 1.175T20 6v15q0 .425-.288.713T19 22q-.425 0-.713-.288T18 21V6q0-.825-.588-1.413T16 4q-.825 0-1.413.588T14 6v12q0 1.65-1.175 2.825T10 22Z"/>`),
		g.Group(children),
	)
}