package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HlsOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.825 15L15.9 13.075q.075-.05.15-.063t.2-.012q.275 0 .45.138t.275.362H19v-1h-2.5q-.425 0-.713-.288T15.5 11.5V10q0-.425.288-.713T16.5 9h3q.425 0 .713.288T20.5 10v.25q0 .325-.213.537T19.75 11q-.25 0-.45-.138t-.25-.362H17v1h2.5q.425 0 .713.288t.287.712V14q0 .425-.288.713T19.5 15h-1.675ZM3.75 15q-.325 0-.538-.213T3 14.25v-4.5q0-.325.213-.537T3.75 9q.325 0 .537.213t.213.537V11h2V9.75q0-.325.213-.537T7.25 9q.325 0 .537.213T8 9.75v4.5q0 .325-.213.537T7.25 15q-.325 0-.537-.213T6.5 14.25V12.5h-2v1.75q0 .325-.213.537T3.75 15Zm16.025 7.6L1.4 4.225L2.8 2.8l18.375 18.4l-1.4 1.4ZM10 15v-3.575l1.5 1.5v.575h.6l1.5 1.5H10Z"/>`),
		g.Group(children),
	)
}