package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JavascriptRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15q-.425 0-.713-.288T12 14v-.25q0-.325.213-.537T12.75 13q.25 0 .438.138t.262.362h2.05v-1H13q-.425 0-.713-.288T12 11.5V10q0-.425.288-.713T13 9h3q.425 0 .713.288T17 10v.25q0 .325-.213.537T16.25 11q-.25 0-.45-.138t-.25-.362H13.5v1H16q.425 0 .713.288T17 12.5V14q0 .425-.288.713T16 15h-3Zm-5.5 0q-.625 0-1.063-.438T6 13.5v-.25q0-.325.213-.537t.537-.213q.325 0 .537.213t.213.537v.25H9V9.75q0-.325.213-.537T9.75 9q.325 0 .537.213t.213.537v3.75q0 .625-.438 1.063T9 15H7.5Z"/>`),
		g.Group(children),
	)
}