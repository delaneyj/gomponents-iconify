package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13l8-5V6l-8 5l-8-5v2l8 5Zm7 7q-.425 0-.713-.288T18 19v-3q0-.425.288-.713T19 15v-1q0-.825.588-1.413T21 12q.825 0 1.413.588T23 14v1q.425 0 .713.288T24 16v3q0 .425-.288.713T23 20h-4Zm1-5h2v-1q0-.425-.288-.713T21 13q-.425 0-.713.288T20 14v1ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v4h-1q-2.075 0-3.538 1.463T16 15v5H4Z"/>`),
		g.Group(children),
	)
}