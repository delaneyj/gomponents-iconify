package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalNotesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16q-1.25 0-2.125-.875T14 13q0-1.25.875-2.125T17 10q1.25 0 2.125.875T20 13q0 1.25-.875 2.125T17 16Zm0-2q.425 0 .713-.288T18 13q0-.425-.288-.713T17 12q-.425 0-.713.288T16 13q0 .425.288.713T17 14Zm-6 9v-4.275q1.05-.6 2.15-1t2.3-.6L17 19l1.55-1.875q1.2.15 2.3.575t2.15 1V23H11Zm1.975-2h3.075l-1.35-1.65q-.45.125-.875.325t-.85.425v.9Zm4.975 0H21v-.9q-.4-.25-.825-.438t-.875-.312L17.95 21Zm-1.9 0Zm1.9 0ZM9 21H3V3h18v7q-.4-.5-.875-.95T19 8.45V5H5v14h4v2ZM7 9h7q.65-.5 1.425-.75T17 8V7H7v2Zm0 4h5q0-.525.112-1.025t.313-.975H7v2Zm0 4h3.15l1.525-.675V15H7v2Zm2 2H5V5h14v3.425q-.45-.2-.963-.313T17 8q-2.075 0-3.538 1.463T12 13q0 .775.213 1.488t.637 1.312L9 17.525V19Zm8-6Z"/>`),
		g.Group(children),
	)
}