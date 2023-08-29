package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluidOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 23q-.825 0-1.413-.588T11 21v-2.075q-2.575-.35-4.288-2.313T5 12V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v9q0 2.65-1.713 4.612T13 18.925V21h5q.425 0 .713.288T19 22q0 .425-.288.713T18 23h-5Zm.75-10h3.15q.05-.25.075-.488T17 12v-1h-3q-.425 0-.713-.288T13 10q0-.425.288-.713T14 9h3V7h-4q-.425 0-.713-.288T12 6q0-.425.288-.713T13 5h4V3H7v8h2.75q.825 0 1.563.375T12.55 12.4q.2.275.525.438t.675.162ZM12 17q1.225 0 2.275-.537T16 15h-2.25q-.825 0-1.563-.375T10.95 13.6q-.2-.275-.525-.437T9.75 13H7.1q.35 1.725 1.713 2.863T12 17Zm-1.05-4Z"/>`),
		g.Group(children),
	)
}