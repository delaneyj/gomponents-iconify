package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RearCameraSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 8q.425 0 .713-.288T17.5 7q0-.425-.288-.713T16.5 6q-.425 0-.713.288T15.5 7q0 .425.288.713T16.5 8ZM13 19h7V5h-7v14ZM2 21v-8h4.2l-1.6 1.6L6 16l4-4l-4-4l-1.4 1.4L6.2 11H2V3h20v18H2Z"/>`),
		g.Group(children),
	)
}