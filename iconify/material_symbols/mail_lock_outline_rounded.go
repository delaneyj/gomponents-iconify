package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailLockOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v4h-2V8l-7.475 4.675q-.125.075-.263.113t-.262.037q-.125 0-.263-.037t-.262-.113L4 8v10h12v2H4Zm8-9l8-5H4l8 5Zm-8 7V8v.25v-1.475v.025V6v.8v-.012V8.25V8v10Zm15 2q-.425 0-.713-.288T18 19v-3q0-.425.288-.713T19 15v-1q0-.825.588-1.413T21 12q.825 0 1.413.588T23 14v1q.425 0 .713.288T24 16v3q0 .425-.288.713T23 20h-4Zm1-5h2v-1q0-.425-.288-.713T21 13q-.425 0-.713.288T20 14v1Z"/>`),
		g.Group(children),
	)
}