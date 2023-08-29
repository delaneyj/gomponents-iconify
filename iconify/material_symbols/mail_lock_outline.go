package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailLockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v3.975h-1L20 10V8l-8 5l-8-5v10h12v2H4Zm8-9l8-5H4l8 5Zm-8 7V6v12Zm15 2q-.425 0-.713-.288T18 19v-3q0-.425.288-.713T19 15v-1q0-.825.588-1.413T21 12q.825 0 1.413.588T23 14v1q.425 0 .713.288T24 16v3q0 .425-.288.713T23 20h-4Zm1-5h2v-1q0-.425-.288-.713T21 13q-.425 0-.713.288T20 14v1Z"/>`),
		g.Group(children),
	)
}