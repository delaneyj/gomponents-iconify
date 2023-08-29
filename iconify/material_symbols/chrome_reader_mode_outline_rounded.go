package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeReaderModeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h7V6H4v12Zm9 0h7V6h-7v12Zm5.25-3q.325 0 .537-.213T19 14.25q0-.325-.213-.537t-.537-.213h-3.5q-.325 0-.537.213T14 14.25q0 .325.213.537t.537.213h3.5Zm0-2.5q.325 0 .537-.213T19 11.75q0-.325-.213-.537T18.25 11h-3.5q-.325 0-.537.213T14 11.75q0 .325.213.537t.537.213h3.5Zm0-2.5q.325 0 .537-.213T19 9.25q0-.325-.213-.537T18.25 8.5h-3.5q-.325 0-.537.213T14 9.25q0 .325.213.537t.537.213h3.5ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}