package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenLockPortraitSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16v-5h1v-1q0-.825.588-1.413T12 8q.825 0 1.413.588T14 10v1h1v5H9Zm2-5h2v-1q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10v1ZM5 23V1h14v22H5Zm2-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}