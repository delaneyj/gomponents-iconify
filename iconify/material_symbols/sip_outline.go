package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SipOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2V9h-2v6Zm3 0h1.5v-2H18q.425 0 .713-.288T19 12v-2q0-.425-.288-.713T18 9h-4v6Zm-9 0h4q.425 0 .713-.288T10 14v-1.75q0-.425-.288-.713T9 11.25H6.5v-.75H10V9H6q-.425 0-.713.288T5 10v1.75q0 .425.288.713T6 12.75h2.5v.75H5V15Zm10.5-3.5v-1h2v1h-2ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}