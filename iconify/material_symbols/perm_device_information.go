package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermDeviceInformation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16.5V11h2v5.5h-2ZM12 9q-.425 0-.713-.288T11 8q0-.425.288-.713T12 7q.425 0 .713.288T13 8q0 .425-.288.713T12 9ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}