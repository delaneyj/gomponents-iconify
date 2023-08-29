package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPostOfficeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21v-8q0-.425.288-.713T3 12h3V8q0-2.5 1.75-4.25T12 2h4q2.5 0 4.25 1.75T22 8v13q0 .425-.288.713T21 22q-.425 0-.713-.288T20 21v-2h-4v2q0 .425-.288.713T15 22H3Zm13-5h4V8q0-1.65-1.175-2.825T16 4h-4q-1.65 0-2.825 1.175T8 8v4h7q.425 0 .713.288T16 13v4Zm-7-.15L14 14H4l5 2.85ZM4 20h10v-4.25l-4 2.275q-.225.125-.475.2T9 18.3q-.275 0-.525-.075t-.475-.2L4 15.75V20Zm0-6v.65v-.025v1.35v-.225V20v-4.25v.225v-1.35v.025V14Zm7-4q-.425 0-.713-.288T10 9q0-.425.288-.713T11 8h6q.425 0 .713.288T18 9q0 .425-.288.713T17 10h-6Z"/>`),
		g.Group(children),
	)
}