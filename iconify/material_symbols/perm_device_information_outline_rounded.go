package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermDeviceInformationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.5q-.425 0-.713-.288T11 15.5V12q0-.425.288-.713T12 11q.425 0 .713.288T13 12v3.5q0 .425-.288.713T12 16.5ZM12 9q-.425 0-.713-.288T11 8q0-.425.288-.713T12 7q.425 0 .713.288T13 8q0 .425-.288.713T12 9ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-3v1h10v-1H7Zm0-2h10V6H7v12ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}