package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCallRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16q.425 0 .713-.288T11 15v-2h2q.425 0 .713-.288T14 12q0-.425-.288-.713T13 11h-2V9q0-.425-.288-.713T10 8q-.425 0-.713.288T9 9v2H7q-.425 0-.713.288T6 12q0 .425.288.713T7 13h2v2q0 .425.288.713T10 16Zm-6 4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l3.15-3.15q.225-.225.537-.113T22 7.7v8.6q0 .35-.313.463t-.537-.113L18 13.5V18q0 .825-.588 1.413T16 20H4Z"/>`),
		g.Group(children),
	)
}