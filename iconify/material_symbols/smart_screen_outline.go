package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartScreenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.25 12.75q-.325 0-.537-.213T12.5 12q0-.325.213-.537t.537-.213q.325 0 .537.213T14 12q0 .325-.213.537t-.537.213Zm-5 0q-.325 0-.537-.213T7.5 12q0-.325.213-.537t.537-.213q.325 0 .537.213T9 12q0 .325-.213.537t-.537.213Zm7.5 0q-.325 0-.537-.213T15 12q0-.325.213-.537t.537-.213q.325 0 .537.213T16.5 12q0 .325-.213.537t-.537.213Zm-5 0q-.325 0-.537-.213T10 12q0-.325.213-.537t.537-.213q.325 0 .537.213T11.5 12q0 .325-.213.537t-.537.213ZM3 19q-.825 0-1.413-.588T1 17V7q0-.825.588-1.413T3 5h18q.825 0 1.413.588T23 7v10q0 .825-.588 1.413T21 19H3Zm1-2V7H3v10h1Zm2 0h12V7H6v10Zm14 0h1V7h-1v10ZM4 7H3h1Zm16 0h1h-1Z"/>`),
		g.Group(children),
	)
}