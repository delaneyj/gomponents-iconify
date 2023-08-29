package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraIndoorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Zm3-4h4q.425 0 .713-.288T14 16v-1l1.275.675q.25.125.488-.025t.237-.425v-2.45q0-.275-.238-.425t-.487-.025L14 13v-1q0-.425-.288-.713T13 11H9q-.425 0-.713.288T8 12v4q0 .425.288.713T9 17Z"/>`),
		g.Group(children),
	)
}