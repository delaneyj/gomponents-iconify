package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicFormOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 11q-.825 0-1.413-.588T2 9V6q0-.825.588-1.413T4 4h8q.425 0 .713.288T13 5v5q0 .425-.288.713T12 11H4Zm0-2h7V6H4v3Zm0 11q-.825 0-1.413-.588T2 18v-3q0-.825.588-1.413T4 13h10q.425 0 .713.288T15 14v5q0 .425-.288.713T14 20H4Zm0-2h9v-3H4v3Zm13-7h-1q-.425 0-.713-.288T15 10V5q0-.425.288-.713T16 4h4.525q.525 0 .825.438t.1.937L20 9h.45q.55 0 .825.463t.075.962l-3.4 9.225q-.075.2-.225.275t-.325.025q-.175-.05-.288-.175T17 19.45V11ZM4 9V6v3Zm0 9v-3v3ZM6.25 7.5q0-.325-.213-.537T5.5 6.75q-.325 0-.537.213T4.75 7.5q0 .325.213.537t.537.213q.325 0 .537-.213T6.25 7.5Zm-.75 9.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T5.5 15.75q-.325 0-.537.213t-.213.537q0 .325.213.537t.537.213Z"/>`),
		g.Group(children),
	)
}