package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportantDevicesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21q-.425 0-.713-.288T16 20v-7.95q0-.425.288-.713T17 11.05h4q.425 0 .713.288t.287.712V20q0 .425-.288.713T21 21h-4Zm0-2h4v-5.95h-4V19Zm-9 2v-2h2v-2H4q-.825 0-1.413-.588T2 15V5q0-.825.588-1.413T4 3h14q.825 0 1.413.588T20 5v4.05h-2V5H4v10h10v2h-2v2h2v2H8Zm.7-7.25L11 12l2.3 1.75l-.85-2.85l2.3-1.85H11.9l-.9-2.8l-.9 2.8H7.25l2.3 1.85l-.85 2.85ZM11 10Z"/>`),
		g.Group(children),
	)
}