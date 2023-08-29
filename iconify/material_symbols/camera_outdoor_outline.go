package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOutdoorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 18q-.425 0-.713-.288T12 17v-4q0-.425.288-.713T13 12h4q.425 0 .713.288T18 13v1l2-1.05v4.1L18 16v1q0 .425-.288.713T17 18h-4Zm-9 3V9l8-6l8 6v2h-2v-1l-6-4.5L6 10v9h14v2H4Zm8-8.75Z"/>`),
		g.Group(children),
	)
}