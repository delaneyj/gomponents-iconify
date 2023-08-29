package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOutdoorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v1h-7q-.825 0-1.413.588T11 13v4q0 .825.588 1.413T13 19h6q.425 0 .713.288T20 20q0 .425-.288.713T19 21H6Zm7-3q-.425 0-.713-.288T12 17v-4q0-.425.288-.713T13 12h4q.425 0 .713.288T18 13v1l1.275-.675q.25-.125.488.025t.237.425v2.45q0 .275-.238.425t-.487.025L18 16v1q0 .425-.288.713T17 18h-4Z"/>`),
		g.Group(children),
	)
}