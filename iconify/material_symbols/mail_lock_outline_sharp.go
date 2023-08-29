package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailLockOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v5.975h-1L20 10V8l-8 5l-8-5v10h12v2H2Zm10-9l8-5H4l8 5Zm-8 7V6v12Zm14 2v-5h1v-1q0-.825.588-1.413T21 12q.825 0 1.413.588T23 14v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T21 13q-.425 0-.713.288T20 14v1Z"/>`),
		g.Group(children),
	)
}