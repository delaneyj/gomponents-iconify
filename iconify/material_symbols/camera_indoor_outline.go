package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraIndoorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h4q.425 0 .713-.288T14 16v-1l2 1.05v-4.1L14 13v-1q0-.425-.288-.712T13 11H9q-.425 0-.713.288T8 12v4q0 .425.288.713T9 17Zm-5 4V9l8-6l8 6v12H4Zm2-2h12v-9l-6-4.5L6 10v9Zm6-6.75Z"/>`),
		g.Group(children),
	)
}