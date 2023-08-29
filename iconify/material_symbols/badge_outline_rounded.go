package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18h6v-.45q0-.425-.238-.788T11.1 16.2q-.5-.225-1.012-.338T9 15.75q-.575 0-1.088.113T6.9 16.2q-.425.2-.663.563T6 17.55V18Zm8.75-1.5h2.5q.325 0 .537-.213T18 15.75q0-.325-.213-.537T17.25 15h-2.5q-.325 0-.537.213T14 15.75q0 .325.213.537t.537.213ZM9 15q.625 0 1.063-.438T10.5 13.5q0-.625-.438-1.063T9 12q-.625 0-1.063.438T7.5 13.5q0 .625.438 1.063T9 15Zm5.75-1.5h2.5q.325 0 .537-.213T18 12.75q0-.325-.213-.537T17.25 12h-2.5q-.325 0-.537.213T14 12.75q0 .325.213.537t.537.213ZM4 22q-.825 0-1.413-.588T2 20V9q0-.825.588-1.413T4 7h5V4q0-.825.588-1.413T11 2h2q.825 0 1.413.588T15 4v3h5q.825 0 1.413.588T22 9v11q0 .825-.588 1.413T20 22H4Zm0-2h16V9h-5q0 .825-.588 1.413T13 11h-2q-.825 0-1.413-.588T9 9H4v11Zm7-11h2V4h-2v5Zm1 5.5Z"/>`),
		g.Group(children),
	)
}