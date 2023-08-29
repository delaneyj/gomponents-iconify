package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedroomParentOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 15.5h11v.75q0 .325.213.537t.537.213q.325 0 .537-.213T19 16.25v-3.1q0-.525-.2-.988t-.55-.812V9q0-.825-.588-1.412T16.25 7H13q-.275 0-.525.075T12 7.3q-.225-.15-.475-.225T11 7H7.75q-.825 0-1.413.588T5.75 9v2.35q-.35.35-.55.813t-.2.987v3.1q0 .325.213.537T5.75 17q.325 0 .537-.213t.213-.537v-.75Zm0-1.5v-1q0-.425.288-.713T7.5 12h9q.425 0 .713.288T17.5 13v1h-11Zm.75-3.5v-2h4v2h-4Zm5.5 0v-2h4v2h-4ZM4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Zm0-2h16V4H4v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}