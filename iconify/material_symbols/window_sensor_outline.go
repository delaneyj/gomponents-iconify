package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowSensorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 9q-.425 0-.713-.288T21 8V4q0-.425.288-.713T22 3q.425 0 .713.288T23 4v4q0 .425-.288.713T22 9ZM3 21V3h16v18H3Zm2-10h5v-1h2v1h5V5H5v6Zm0 8h12v-6H5v6Zm0 0h12H5Z"/>`),
		g.Group(children),
	)
}