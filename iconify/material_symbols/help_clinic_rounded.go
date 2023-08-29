package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpClinicRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q.425 0 .713-.288T13 16v-3q0-.425-.288-.713T12 12q-.425 0-.713.288T11 13v3q0 .425.288.713T12 17Zm0-7q.425 0 .713-.288T13 9q0-.425-.288-.713T12 8q-.425 0-.713.288T11 9q0 .425.288.713T12 10ZM6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Z"/>`),
		g.Group(children),
	)
}