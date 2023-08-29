package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntegrationInstructionsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.85 12l1.45-1.45q.3-.3.288-.7t-.288-.7q-.3-.3-.713-.313t-.712.288L6.7 11.3q-.3.3-.3.7t.3.7l2.175 2.175q.3.3.713.288t.712-.313q.275-.3.288-.7t-.288-.7L8.85 12Zm6.3 0l-1.45 1.45q-.3.3-.288.7t.288.7q.3.3.713.312t.712-.287L17.3 12.7q.3-.3.3-.7t-.3-.7l-2.175-2.175q-.3-.3-.713-.288t-.712.313q-.275.3-.287.7t.287.7L15.15 12ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h4.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H19q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7-16.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25Z"/>`),
		g.Group(children),
	)
}