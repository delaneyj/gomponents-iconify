package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberPinRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 13H9q.425 0 .713-.288T10 12v-2q0-.425-.288-.713T9 9H5.5q-.2 0-.35.15T5 9.5v4.75q0 .325.213.537T5.75 15q.325 0 .537-.213t.213-.537V13ZM12 9q-.325 0-.537.213t-.213.537v4.5q0 .325.213.537T12 15q.325 0 .537-.213t.213-.537v-4.5q0-.325-.213-.537T12 9Zm3.25 2.5l2.375 3.25q.075.125.2.188T18.1 15h.3q.25 0 .425-.175T19 14.4V9.625q0-.275-.175-.45T18.375 9q-.275 0-.45.175t-.175.45V12.5l-2.325-3.25q-.1-.125-.225-.188T14.925 9h-.3q-.275 0-.45.175t-.175.45v4.75q0 .275.175.45t.45.175q.275 0 .45-.175t.175-.45V11.5Zm-8.75 0v-1h2v1h-2ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}