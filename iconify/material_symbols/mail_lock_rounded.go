package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailLockRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11L5.3 6.8q-.425-.275-.862-.025T4 7.525q0 .225.1.413t.3.312l7.075 4.425q.25.15.525.15t.525-.15L19.6 8.25q.2-.125.3-.313t.1-.412q0-.5-.437-.75T18.7 6.8L12 11Zm7 9q-.425 0-.713-.288T18 19v-3q0-.425.288-.713T19 15v-1q0-.825.588-1.413T21 12q.825 0 1.413.588T23 14v1q.425 0 .713.288T24 16v3q0 .425-.288.713T23 20h-4Zm1-5h2v-1q0-.425-.288-.713T21 13q-.425 0-.713.288T20 14v1ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v4h-1q-2.075 0-3.538 1.463T16 15v5H4Z"/>`),
		g.Group(children),
	)
}