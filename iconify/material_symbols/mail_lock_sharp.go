package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailLockSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13l8-5V6l-8 5l-8-5v2l8 5ZM2 20V4h20v6h-1q-2.075 0-3.538 1.463T16 15v5H2Zm16 0v-5h1v-1q0-.825.588-1.413T21 12q.825 0 1.413.588T23 14v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T21 13q-.425 0-.713.288T20 14v1Z"/>`),
		g.Group(children),
	)
}