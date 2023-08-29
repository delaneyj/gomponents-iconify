package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDeleteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h5.175q.4 0 .763.15t.637.425L12 6h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm10.5-3h2q.625 0 1.063-.438T18 15.5v-4h.25q.325 0 .537-.213T19 10.75q0-.325-.213-.537T18.25 10H16.5v-.5q0-.2-.15-.35T16 9h-1q-.2 0-.35.15t-.15.35v.5h-1.75q-.325 0-.537.213T12 10.75q0 .325.213.537t.537.213H13v4q0 .625.438 1.063T14.5 17Zm0-5.5h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}